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

	// create a struct that stores user input
	type UserSignUp struct {
		UserName string `json:"userName" form:"userName"`
		Email string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
		Phone string `json:"phone" form:"phone"`
		Age int `json:"age" form:"age"`
		Language string `json:"language" form:"language"`
	}

	// create a var and give type as struct
	var inputUser UserSignUp

	// bind it into json first 
	if err := c.ShouldBindJSON(&inputUser); err != nil {
       c.JSON(400, gin.H{
		"msg": "invalid req",
	   })
	   return
	}

	// validations 
	if inputUser.UserName == ""|| inputUser.Email == "" || inputUser.Password == "" || inputUser.Phone == "" || inputUser.Age == 0 || inputUser.Language == "" {
		c.JSON(400, gin.H{
			"msg": "fill all fields",
		})
		return
	}
	if !strings.Contains(inputUser.Email, "@"){
		c.JSON(400, gin.H{
			"msg": "Invalid email",
		})
		return
	}
	if len(inputUser.Password) < 6 {
		c.JSON(400, gin.H{
			"msg": "atleast 6",
		})
		return
	}
	if len(inputUser.Phone) < 10 {
		c.JSON(400, gin.H{
			"msg": "must be atleast 10",
		})
		return
	}

	// check duplicate
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": inputUser.Email})
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid",
		})
		return
	}
	if count > 0 {
		c.JSON(400, gin.H{
			"msg": "user already exists",
		})
		return
	}

	// hash the pass 
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(inputUser.Password), 10)
	emailToken := generateToken(8)
	phoneToken := generateToken(8)

	// create model var
	var user models.User

	user.UserName = inputUser.UserName
    user.Email = inputUser.Email
	user.Password = string(hashPass)
	user.Age = inputUser.Age
	user.Phone = inputUser.Phone
	user.Language = inputUser.Language
	user.UserVerified.Email = false
	user.UserVerifyToken.Email = emailToken
	user.UserVerifyToken.Phone = phoneToken

	// send email for verify 
	go func() {
		emailData := utils.EmailData{
			From: "Tea, insa",
			To: inputUser.Email,
			Html: fmt.Sprintf(`<a href="%s/api/public/user/emailverify/%s"></a>`, config.AppConfig.URL, emailToken),
		}
		_  = utils.SendEmail(emailData); 
	}()

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "err",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "user created bro verify ur email",
	})
}

func EmailVerify(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	token := c.Param("token")

	// find usre with this tokem
	var user models.User

	err := userCollection.FindOne(ctx, bson.M{"userVerifyToken.email": token}).Decode(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid token",
		})
		return
	}

	if user.UserVerified.Email {
		c.JSON(200, gin.H{
			"msg": "user already verified",
		})
		return
	}

	// update 
	update := bson.M{
		"$set": bson.M{
			"userVerfied.email": true,
			"userVerifyToken": nil,
			"updated_at": time.Now(),
		}}


	_, err = userCollection.UpdateByID(ctx, user.ID, update)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "error while uodating",
		})
		return
	}	
	c.JSON(200, gin.H{
		"msg": "email verified✅",
	})
}