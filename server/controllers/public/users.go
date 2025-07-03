package public

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"github.com/AbdulRahman-04/Go_Backend_Practice/config"
	"github.com/AbdulRahman-04/Go_Backend_Practice/models"
	"github.com/AbdulRahman-04/Go_Backend_Practice/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func UserCollect() {
	userCollection = utils.MongoClient.Database("Insta_Backend").Collection("user")
}

var JwtKey = []byte(config.AppConfig.JWTKEY)
var URL = config.AppConfig.URL

func userGenerateToken(length int) string {
	d := make([]byte, length)
	_, _ = rand.Read(d)
	return hex.EncodeToString(d)
}

// user signup api
func UserSignUp(c *gin.Context) {
	// ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// type strutc bnao and bind it into json
	type UserSignUp struct {
		UserName string `json:"username" form:"username"`
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
		Phone    string `json:"phone" form:"phone"`
		Age      int    `json:"age" form:"age"`
		Language string `json:"language" form:"language"`
	}

	// create a var and bind into it
	var inputUser UserSignUp

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validations
	if inputUser.UserName == "" || inputUser.Email == "" || inputUser.Password == "" || inputUser.Age == 0 || inputUser.Phone == "" || inputUser.Language == "" {
		c.JSON(400, gin.H{
			"msg": "invalid fill all fields",
		})
		return
	}

	if !strings.Contains(inputUser.Email, "@") {
		c.JSON(400, gin.H{
			"msg": "invalid email",
		})
		return
	}

	if len(inputUser.Password) < 6 {
		c.JSON(400, gin.H{
			"msg": "invalid pass lengt",
		})
		return
	}

	if len(inputUser.Phone) < 10 {
		c.JSON(400, gin.H{
			"msg": "invalid length number",
		})
		return
	}

	// duplicate check
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": inputUser.Email})
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db error",
		})
		return
	}

	if count > 0 {
		c.JSON(400, gin.H{
			"msg": "invalid user already exists",
		})
		return
	}

	// hash pass
	hashPass, err := bcrypt.GenerateFromPassword([]byte(inputUser.Password), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldnt hash pass",
		})
		return
	}

	emailToken := userGenerateToken(8)
	phoneToken := userGenerateToken(8)

	// create a var and put all values inside it
	var user models.User

	user.UserName = inputUser.UserName
	user.Email = inputUser.Email
	user.Password = string(hashPass)
	user.Phone = inputUser.Phone
	user.Role = "user"
	user.Age = inputUser.Age
	user.Language = inputUser.Language
	user.UserVerified.Email = false
	user.UserVerifyToken.Email = emailToken
	user.UserVerifyToken.Phone = phoneToken
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// send email func
	go func() {

		emailData := utils.EmailData{
			From:    "team insta",
			To:      inputUser.Email,
			Subject: "email verify",
		    Html: fmt.Sprintf(`<a href="%s/api/public/user/emailverify/%s">Click here to verify your email</a>`, URL, emailToken),	
		}

		_ = utils.SendEmail(emailData)

	}()

	// add the obj into db
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldn't add user to db",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "user added successfully!✅",
	})

}

// email verify
func UserEmailVerify(c *gin.Context) {
	// ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token := c.Param("token")

	// compare token in db
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"userVerifyToken.emailToken": token}).Decode(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid token",
		})
		return
	}

	// check if link hasnt clicked more than once
	if user.UserVerified.Email {
		c.JSON(200, gin.H{
			"msg": "email already verified",
		})
		return
	}

	// update into db
	update := bson.M{
		"$set": bson.M{
			"userVerified.emailVerified": true,
			"userVerifyToken.emailToken": nil,
			"updated_at":                 time.Now(),
		}}

	// update the changes into dv
	_, err = userCollection.UpdateByID(ctx, user.ID, update)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "db error",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "email verified✅",
	})

}

// user signin
func UserSignIn(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// type struct
	type UserSignIn struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	var inputUser UserSignIn

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validations
	if !strings.Contains(inputUser.Email, "@") || len(inputUser.Password) < 6 || inputUser.Email == "" || inputUser.Password == "" {
		c.JSON(400, gin.H{
			"msg": "pls fill valid email and password",
		})
		return
	}

	// check if email exist in db
	var user models.User

	err := userCollection.FindOne(ctx, bson.M{"email": inputUser.Email}).Decode(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid email not found",
		})
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"msg": " invalid password ",
		})
		return
	}

	// token generate
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"role":  user.Role,
		"email": user.Email,
		"exp":   time.Now().Add(5 * time.Hour).Unix(),
	}).SignedString(JwtKey)

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldn't generate token",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "logged in successfully✅", "token": token})

}

// change pass api 
func ChangeUserPass(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// type struct 
	type ChangeUserPass struct {
		Email string         `json:"email" form:"email"`
		OldPassword string   `json:"oldpassword" form:"oldpassword"` 
		NewPassword string   `json:"newpassword" form:"newpassword"`
	}

	// bind into json
	var inputUser ChangeUserPass

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validations 
	if !strings.Contains(inputUser.Email, "@") || len(inputUser.NewPassword) < 6 || inputUser.Email == "" || inputUser.NewPassword == "" || inputUser.OldPassword == "" {
		c.JSON(400, gin.H{
			"msg": "invalid pls fill all fields correctly",
		})
		return
	}

	// find user email in db 
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": inputUser.Email}).Decode(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid no user email found",
		})
		return
	}

	// compare old password 
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.OldPassword))
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid password",
		})
		return
	}

	// hash the new password
	hashPass , err := bcrypt.GenerateFromPassword([]byte(inputUser.NewPassword), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldnt hash pass",
		})
		return
	}

	// update the pass 
	update := bson.M{
		"$set": bson.M{
			"password": string(hashPass),
			"updated_at": time.Now(),
		}}

	// update change into db 
	_, err = userCollection.UpdateByID(ctx, user.ID, update)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db error",
		})
		return
	} 

	c.JSON(200, gin.H{
			"msg": "password changed✅",
		})
		
}

// forgot password api 
func ForgotPassUser(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// type struct 
	type ForgotPassUser struct {
		Email string `json:"email" form:"email"`
	}

	var inputUser ForgotPassUser

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validation 
	if !strings.Contains(inputUser.Email, "@") || inputUser.Email == "" {
		c.JSON(400, gin.H{
			"msg": "pls type an email",
		})
		return
	}

	// find the email in db 
	var user models.User

	err := userCollection.FindOne(ctx, bson.M{"email": inputUser.Email}).Decode(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid user",
		})
		return
	}

	// temp pass generate 
	tempPass := userGenerateToken(8)
	hashPass, err := bcrypt.GenerateFromPassword([]byte(tempPass), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldnt hash pass",
		})
		return
	}

	// update into db 
	update := bson.M{
		"$set": bson.M{
			"password": string(hashPass),
			"updated_at": time.Now(),
		}}
	// update into db 
	_, err = userCollection.UpdateByID(ctx, user.ID, update)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	} 

	// send pass on email
	go func(){
 
		emailData := utils.EmailData {
			From: "team insta",
			To: inputUser.Email,
			Subject: "temporary pass",
            Html: fmt.Sprintf(`<h2>user %s ur password is <strong>%s</strong></h2>`, user.UserName, tempPass),
		}
      _ = utils.SendEmail(emailData)
	}()	

	c.JSON(200, gin.H{
			"msg": "temporary email sent✅",
		})
		
}