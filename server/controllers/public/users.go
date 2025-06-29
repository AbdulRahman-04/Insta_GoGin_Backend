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
	_, _ = rand.Read(d)
	return hex.EncodeToString(d)
}

func UserSignUp(c *gin.Context){
	// create ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create a struct for storing user input inside it
	type UserSignUp struct {
		UserName string `json:"userName" form:"userName"`
		Email string `form:"email" json:"email"`
		Password string `json:"password" form:"password"`
		Phone string `json:"phone" form:"phone"`
		Age int `json:"age" form:"age"`
		Language string `json:"language" form:"language"`
	}

	// store thsi struct in var and bind into json
	var inputUser UserSignUp

	if err := c.ShouldBindJSON(&inputUser); err != nil {
       c.JSON(400, gin.H{
		"msg": "Invalid request",
	   })
	   return
	}

	// validations 
	if inputUser.UserName == "" || inputUser.Email == "" || inputUser.Password == "" || inputUser.Phone == "" || inputUser.Age == 0 || inputUser.Language == "" {
		c.JSON(400, gin.H{
			"msg": "Please input all fields",
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
			"msg": "6 digits toh rkh",
		})
		return
	}
	if len(inputUser.Phone) < 10 {
		c.JSON(400, gin.H{
			"msg": "must be 10",
		})
		return
	}

	// check if user exusts in db
	count, err := userCollection.CountDocuments(ctx, bson.M{"email" : inputUser.Email})
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "db error",
		})
		return
	}

	if count > 0 {
		c.JSON(404, gin.H{
			"msg": "user already exists",
		})
		return
	}

	// hash the password 
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(inputUser.Password), 10)
	emailToken := generateToken(8)
	phoneToken := generateToken(8)

	// create a var for pushing into db
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

	// send email func 
	go func() {
		emailData := utils.EmailData{
			From: "team insta",
			To: inputUser.Email,
			Subject: "email verify",
			Html: fmt.Sprintf(`<a href="%s/api/public/user/emailverify/%s"></a>`, config.AppConfig.URL, emailToken ),
		}
		_ = utils.SendEmail(emailData)
	}()

	_, err = userCollection.InsertOne(ctx, user)

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "Error while creating user",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "user created✅, verify ur email",
	})

}