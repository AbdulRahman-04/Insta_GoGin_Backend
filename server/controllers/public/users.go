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
	_,_ = rand.Read(d)
	return hex.EncodeToString(d)
}

func UserSignUp(c *gin.Context){
	// create ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// type struct
	type UserSignUp struct {
		UserName string `json:"userName" form:"userName"`
		Email string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
		Phone string `json:"phone" form:"phone"`
		Age int `json:"age" form:"age"`
		Language string `json:"language" form:"language"`
	}

	// var and bind
	var inputUser UserSignUp

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validations
	if inputUser.UserName == "" || inputUser.Email == "" || inputUser.Password == "" || inputUser.Age == 0 || inputUser.Phone == "" || inputUser.Language == ""{
		c.JSON(400, gin.H{
			"msg": "Please enter all foeld",
		})
		return
	}
	if !strings.Contains(inputUser.Email, "@"){
		c.JSON(400, gin.H{
			"msg": "enter valid email",
		})
		return
	}
	if len(inputUser.Password) < 6 {
		c.JSON(400, gin.H{
			"msg": "enter 6 digit pass",
		})
		return
	}
	if len(inputUser.Phone) < 10 {
		c.JSON(400, gin.H{
			"msg": "enter 10 digits number",
		})
		return
	}

	// duplicate check 
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": inputUser.Email})
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "db error",
		})
		return
	}

	if count > 0 {
		c.JSON(400, gin.H{
			"msg": "already exists",
		})
		return
	}

	// hash pass nd email tokens
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(inputUser.Password), 10)
	emailToken := generateToken(8)
	phoneToken := generateToken(8)

	// create obj var to push into db
	var user models.User

	user.UserName = inputUser.UserName
	user.Email = inputUser.Email
	user.Password = string(hashPass)
	user.Age = inputUser.Age
	user.Language = inputUser.Language
	user.Phone = inputUser.Phone
	user.UserVerified.Email = false
	user.UserVerifyToken.Email = emailToken
	user.UserVerifyToken.Phone = phoneToken

	// send email func 
	go func(){
		emailData := utils.EmailData{
			From: "Team insta",
			To: inputUser.Email,
			Subject: "email verify",
			Html: fmt.Sprintf(`<a href="%s/api/public/user/emailverify/%s"></a>`, config.AppConfig.URL, emailToken),
		}
		_ = utils.SendEmail(emailData)
	} ()

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldn't register user",
		})
		return
	}

	c.JSON(200, gin.H{
			"msg": "user registered, please verify email",
		})
		return
}

func EmailVerify(c *gin.Context){
	token := c.Param("token")
	// ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	

	// find one 
	var user models.User

	err := userCollection.FindOne(ctx, bson.M{"userVerifyToken.email": token}).Decode(&user)
	if err != nil {
		 c.JSON(400, gin.H{
			"msg": "invalid token",
		 })
		 return
	}

	// check if already updated
	if user.UserVerified.Email {
		c.JSON(200, gin.H{
			"msg": "already verified",
		 })
		 return
	}

	// update the coontent
	update := bson.M{
		"$set": bson.M{
			"userVerified.email": true,
			"userVerifyToken.email": nil,
			"updated_at": time.Now(),
		}}
    
	// update the valuesi n db 
	_, err = userCollection.UpdateByID(ctx, user.ID, update)
	
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "error updating",
		 })
		 return
	}

	c.JSON(200, gin.H{
			"msg": "email verified success✅",
		 })
		 return

}