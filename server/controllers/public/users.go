package public

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/AbdulRahman-04/Go_Backend_Practice/config"
	"github.com/AbdulRahman-04/Go_Backend_Practice/models"
	"github.com/AbdulRahman-04/Go_Backend_Practice/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func userCollect() {
	userCollection = utils.MongoClient.Database("GO_BACKEND_practice").Collection("user")
}

var JWTKEY = []byte(config.AppConfig.JWTKEY)
var Url = config.AppConfig.URL

func generateUserToken(length int) string {
	d := make([]byte, length)
	_, _ = rand.Read(d)
	return hex.EncodeToString(d)
}

// User signup api
func UserSignup(c *gin.Context) {
	// ctx, cancel
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// type struc to store user input
	type UserSignup struct {
		UserName string `json:"username" form:"username"`
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
		Phone    string `json:"phone" form:"phone"`
		Age      int    `json:"age" form:"age"`
		Language string `json:"language" form:"language"`
	}

	// put this struct into var and bind it
	var inputUser UserSignup

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(400, gin.H{
			"MSG": "Invalid request",
		})
		return
	}

	// validations
	if inputUser.UserName == "" || inputUser.Email == "" || inputUser.Password == "" || inputUser.Phone == "" || inputUser.Age == 0 || inputUser.Language == "" {
		c.JSON(400, gin.H{
			"MSG": "Pls fill all fields",
		})
		return
	}

	if !strings.Contains(inputUser.Email, "@") {
		c.JSON(400, gin.H{
			"MSG": "Invalid email",
		})
		return
	}

	if len(inputUser.Password) < 6 {
		c.JSON(400, gin.H{
			"MSG": "Invalid password",
		})
		return
	}
	if len(inputUser.Phone) < 10 {
		c.JSON(400, gin.H{
			"MSG": "Invalid number",
		})
		return
	}

	// duplicate check
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": inputUser.Email})
	if err != nil {
		c.JSON(400, gin.H{
			"MSG": "Invalid db error",
		})
		return
	}

	if count > 0 {
		c.JSON(400, gin.H{
			"MSG": "user already exists",
		})
		return
	}

	// hash the pass
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(inputUser.Password), 10)
	emailToken := generateUserToken(8)
	phoneToken := generateUserToken(8)

	// create user model var
	var user models.User

	user.UserName = inputUser.UserName
	user.Email = inputUser.Email
	user.Password = string(hashPass)
	user.Phone = inputUser.Phone
	user.Age = inputUser.Age
	user.Role = "user"
	user.Language = inputUser.Language
	user.UserVerified.Email = false
	user.UserVerifyToken.Email = emailToken
	user.UserVerifyToken.Phone = phoneToken

	// send email for verify
	go func() {

		emailData := utils.EmailData{
			From:    "Team insta",
			To:      inputUser.Email,
			Subject: "email verify",
			Html:    fmt.Sprintf(`<a href="%s/api/public/user/emailverify/%s"></a>`, Url, emailToken),
		}

		_ = utils.SendEmail(emailData)
	}()

	// insert the user into db
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(400, gin.H{
			"MSG": "Invalid db error, couldn't add user",
		})
		return
	}

	c.JSON(200, gin.H{
		"MSG": "User registered, please verify ur email and login",
	})
	
}

// email verify
func EmailVerify(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// token from param
	token := c.Param("token")

	// find the matchin token in db
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"userVerifyToken.emailToken": token}).Decode(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"MSG": "Invalid token",
		})
		return
	}

	// check if linked hasnt been clicked more than once
	if user.UserVerified.Email {
		c.JSON(400, gin.H{
			"MSG": "email already verified",
		})
		return
	}

	// update
	update := bson.M{
		"$set": bson.M{
			"userVerified.emailVerified": true,
			"userVerifyToken.emailToken": nil,
			"updated_at":                 time.Now(),
		}}

	// update the changes into db
	_, err = userCollection.UpdateByID(ctx, user.ID, update)
	if err != nil {
		c.JSON(400, gin.H{
			"MSG": "couldn't update changes",
		})
		return
	}

	c.JSON(200, gin.H{
		"MSG": "email verified bro✅",
	})
	
}

// user login api
func UserSignin(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// type struct
	type UserSignin struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	// store into var and bind it
	var inputUser UserSignin

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(400, gin.H{
			"MSG": "invalid request",
		})
		return
	}

	// validations
	if inputUser.Email == "" || inputUser.Password == "" {
		c.JSON(400, gin.H{
			"MSG": "fill all fields",
		})
		return
	}

	if !strings.Contains(inputUser.Email, "@") {
		c.JSON(400, gin.H{
			"MSG": "invalid email",
		})
		return
	}

	if len(inputUser.Password) < 6 {
		c.JSON(400, gin.H{
			"MSG": "invalid password length",
		})
		return
	}

	// find user var models.user
	var user models.User

	err := userCollection.FindOne(ctx, bson.M{"email": inputUser.Email}).Decode(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"MSG": "invalid db error",
		})
		return
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password))

	if err != nil {
		c.JSON(400, gin.H{
			"MSG": "invalid password",
		})
		return
	}

	// jwt token generate
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"role":  user.Role,
		"email": user.Email,
		"exp":   time.Now().Add(5 * time.Second).Unix(),
	}).SignedString(JWTKEY)

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldn't generate token",
		})
	}

	c.JSON(400, gin.H{
		"MSG": "logged in successfullly!", "token": token})
}

// change password
func ChangePassUser(c *gin.Context) {
	// ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// type struct
	type ChangePassUser struct {
		Email       string `json:"email" form:"email"`
		OldPassword string `json:"oldpassword" form:"oldpassword"`
		NewPassword string `json:"newpassword" form:"newpassword"`
	}

	// store into var and bind it
	var inputUser ChangePassUser

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	//    validations

	if inputUser.Email == "" || inputUser.OldPassword == "" || inputUser.NewPassword == "" {
		c.JSON(400, gin.H{
			"msg": "pls fill all fields",
		})
		return
	}

	if !strings.Contains(inputUser.Email, "@") {
		c.JSON(400, gin.H{
			"msg": "invalid email",
		})
		return
	}

	if len(inputUser.NewPassword) < 6 {
		c.JSON(400, gin.H{
			"msg": "invalid length pass",
		})
		return
	}

	// find user in db
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": inputUser.Email}).Decode(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "email not found",
		})
		return
	}

	//    compare old pass
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.OldPassword))
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid old password",
		})
		return
	}

	// hash the new pass
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(inputUser.NewPassword), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldn't hash new pass",
		})
		return
	}

	// update in db
	update := bson.M{
		"$set": bson.M{
			"password":   string(hashedPass), // ✅ Hashed password stored
			"updated_at": time.Now(),
		},
	}

	// update the db
	_, err = userCollection.UpdateByID(ctx, user.ID, update)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "db error",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "user password updated✅",
	})
	

}
