package public

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
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

func Collection(){
	userCollection = utils.MongoClient.Database("Go_PRACTICE_BACKEND").Collection("users")
}

var JWTKEY = []byte(config.AppConfig.JWTKEY)

func generateToken(length int) string {
	d := make([]byte, length)
	_, _ = rand.Read(d)
	return hex.EncodeToString(d)
}

func UserSignup(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// type struct 
	type UserSignup struct {
		UserName string `form:"userName" json:"userName"`
		Email    string `form:"email" json:"email"`
		Password string `form:"password" json:"password"`
		Phone    string `form:"phone" json:"phone"`
		Age      int `form:"age" json:"age"`
		Language string `form:"language" json:"language"`
	}

	// var m store krdo 
	var inputUser UserSignup

	// bind krdo 
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validations
	if inputUser.UserName == "" || inputUser.Email == "" || inputUser.Password == "" || inputUser.Age == 0 || inputUser.Language == "" || inputUser.Phone == "" {
			c.JSON(400, gin.H{
			"msg": "fill all fields",
		})
		return
	}
 
	if !strings.Contains(inputUser.Email, "@"){
		c.JSON(400, gin.H{
			"msg": "invalid email",
		})
		return
	}

	if len(inputUser.Password) < 6 {
		c.JSON(400, gin.H{
			"msg": "must be 6",
		})
		return
	}

	if len(inputUser.Phone) < 10 {
		c.JSON(400, gin.H{
			"msg": "must be 10",
		})
		return
	}

	// check duplicate 
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": inputUser.Email})
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "db err",
		})
		return
	}

	if count > 0 {
		c.JSON(400, gin.H{
			"msg": "user already exists go login",
		})
		return
	}

	// hash the pass 
	hashPass,_ := bcrypt.GenerateFromPassword([]byte(inputUser.Password), 10)
	emailToken := generateToken(8)
	phoneToken := generateToken(8) 

	// create model user var 
	var user models.User

	user.UserName = inputUser.UserName
	user.Email = inputUser.Email
	user.Password = string(hashPass)
	user.Phone = inputUser.Phone
	user.Age = inputUser.Age
	user.Language = inputUser.Language
	user.UserVerified.Email = false
	user.UserVerifyToken.Email = emailToken
	user.UserVerifyToken.Phone = phoneToken

	// send email 
	go func () {
		emailData := utils.EmailData {
			From: "team insta",
			To: inputUser.Email,
			Subject: "email verify",
			Html: fmt.Sprintf(`<a href="%s/api/public/user/emailverify/%s"></a>`, config.AppConfig.URL, emailToken),
		}
       _ = utils.SendEmail(emailData)
	}()

	// add obj into db 
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldn't register",
		})
		return
	}

	c.JSON(200, gin.H{
			"msg": "user registered, verify email first",
		})

}

func EmailVerify(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token := c.Param("token")

	// find user 
	var user models.User

	err := userCollection.FindOne(ctx, bson.M{"userVerifyToken.email": token}).Decode(&user)
	if  err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid toke n",
		})
		return
	}

	// check if already verified or not
	if user.UserVerified.Email {
		c.JSON(200, gin.H{
			"msg": "already verified✅",
		})
		return
	}

	// update 
	update := bson.M{
		"$set": bson.M{
			"userVerified.email": true,
			"userVerifyToken.email": nil,
			"updated_at": time.Now(),
		}}

	// finally update the changes in db 
	_, err = userCollection.UpdateByID(ctx, user.ID, update)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldnt update",
		})
		return
	} 

	c.JSON(200, gin.H{
			"msg": "email verified✅",
		})
		return
}

func UserSignIn(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// type  struct 
	type UserSignIn struct {
		Email string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	// store into var and bind to json
	var inputUser UserSignIn

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validations 
	if inputUser.Email == "" || inputUser.Password == "" {
		c.JSON(400, gin.H{
			"msg": "fill all fields",
		})
		return
	}

	if !strings.Contains(inputUser.Email, "@"){
		c.JSON(400, gin.H{
			"msg": "invalid email",
		})
		return
	}

	if len(inputUser.Password) < 6 {
		c.JSON(400, gin.H{
			"msg": "invalid lenght of password",
		})
		return
	}

	// find the user with email in db 
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": inputUser.Email}).Decode(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "user doesn't exists",
		})
		return
	}

	// compare pass 
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid pass",
		})
		return
	}

	// jwt token generate 
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"role": user.Role,
		"email": user.Email,
		"exp": time.Now().Add(5*time.Hour).Unix(),
	}).SignedString(JWTKEY)

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "token generation failed",
		})
		return
	}

	c.JSON(400, gin.H{
			"msg": "logged in bro","token": token})
		
}