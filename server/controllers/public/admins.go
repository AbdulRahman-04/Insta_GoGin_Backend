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

var adminCollection *mongo.Collection

func admCollection(){
	adminCollection = utils.MongoClient.Database("Go_PRACTICE_BACKEND").Collection("admin")
}

var jwtKey = []byte(config.AppConfig.JWTKEY)

func generateAdminToken(length int) string {
	d := make([]byte, length)
	_,  _ = rand.Read(d)
	return hex.EncodeToString(d)
}

func AdminSignup(c *gin.Context){
	// ctx 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create type struct 
	type AdminSignup struct {
		AdminName string `json:"adminName" form:"adminName"`
		Email string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
		Phone string `json:"phone" form:"phone"`
		Age int `json:"age" form:"age"`
		Language string `json:"language" form:"language"`
	}

	// assign struct to var
	var inputAdmin AdminSignup

	// bind into json
	if err := c.ShouldBindJSON(&inputAdmin); err != nil {
		c.JSON(400, gin.H{
			"msg": "Invalid Request",
		})
		return
	}

	// validations 
	if inputAdmin.AdminName == "" || inputAdmin.Email == "" || inputAdmin.Password == "" || inputAdmin.Age == 0 || inputAdmin.Phone == "" || inputAdmin.Language == "" {
		c.JSON(400, gin.H{
			"msg": "fill all fields",
		})
		return
	}

	if!strings.Contains(inputAdmin.Email, "@"){
		c.JSON(400, gin.H{
			"msg": "Invalid email",
		})
		return
	}

	if len(inputAdmin.Password) < 6 {
		c.JSON(400, gin.H{
			"msg": "Invalid length pass",
		})
		return
	}

	if len(inputAdmin.Phone) < 10 {
		c.JSON(400, gin.H{
			"msg": "Invalid mobile no length",
		})
		return
	}

	// duplicate check 
	count, err := adminCollection.CountDocuments(ctx, bson.M{"email": inputAdmin.Email})
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "db error",
		})
		return
	}
	if count > 0 {
		c.JSON(400, gin.H{
			"msg": "admin already exists",
		})
		return
	}

	//HASH PASS ND EM TOKEN
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(inputAdmin.Password), 10)
	emailToken := generateAdminToken(8)
	phoneToken := generateAdminToken(8)

	// create model var admin
	var admin models.Admin

	admin.AdminName = inputAdmin.AdminName
	admin.Email = inputAdmin.Email
	admin.Password = string(hashPass)
	admin.Age = inputAdmin.Age
	admin.Phone = inputAdmin.Phone
	admin.Role = "admin"
	admin.Language = inputAdmin.Language
	admin.AdminVerified.Email = false
	admin.AdminVerifyToken.Email = emailToken
	admin.AdminVerifyToken.Phone = phoneToken
	admin.CreatedAt = time.Now()
	admin.UpdatedAt = time.Now()

	// send email 
	go func () {
		emailData := utils.EmailData {
			From: "Team insta private",
			To: inputAdmin.Email,
			Subject: "Email verify",
			Html: fmt.Sprintf(`<a href="%s/api/public/admin/%s"></a>`, config.AppConfig.URL, emailToken),
		}

		_ = utils.SendEmail(emailData)
	}()

	// now save the data into db 
	_, err = adminCollection.InsertOne(ctx, admin)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "unable to add admin",
		})
		return
	}

	c.JSON(200, gin.H{
			"msg": "Admin Added✅, please login sir!",
		})


}

func EmailVerifyAdmin(c * gin.Context){
	// ctx 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// take token from param
	token := c.Param("token")

	// compare the token in db 
	var admin models.Admin

	err := adminCollection.FindOne(ctx, bson.M{"adminVerifyToken.email": token}).Decode(&admin)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "db error",
		})
		return
	}

	// check if link has clicked more than once 
	if admin.AdminVerified.Email {
		c.JSON(400, gin.H{
			"msg": "admin email verified already bro",
		})
		return
	}

	// update the data first 
	update := bson.M{
		"$set": bson.M{
			"adminVerified.email": true,
			"adminVerifyToken": nil,
			"updated_at": time.Now(),
		}}

	// update the data into db 
	_, err = adminCollection.UpdateByID(ctx, admin.ID, update) 
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldn't update admin",
		})
		return
	}

	c.JSON(400, gin.H{
			"msg": "Admin email verified bro✅",})
		

}

func AdminSignin(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	type AdminSignin struct {
		Email string `json:"email" bson:"email"`
		Password string `json:"password" form:"password"`
	}

	// bind it into a var
	var inputAdmin AdminSignin

	if err := c.ShouldBindJSON(&inputAdmin); err != nil {
		c.JSON(400, gin.H{
			"msg": "Invalid request",
		})
		return
	}

	// validations 
	if inputAdmin.Email == "" || inputAdmin.Password == "" {
		c.JSON(400, gin.H{
			"msg": "invalid email or pass",
		})
		return
	}

	if !strings.Contains(inputAdmin.Email, "@"){
		c.JSON(400, gin.H{
			"msg": "Invalid email",
		})
		return
	}
	if len(inputAdmin.Password) < 6 {
		c.JSON(400, gin.H{
			"msg": "Invalid password",
		})
		return
	}

	// find the email in db 
	var admin models.Admin
	err := adminCollection.FindOne(ctx, bson.M{"email": inputAdmin.Email}).Decode(&admin)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "db error",
		})
		return
	}

	// compare password 
    err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(inputAdmin.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "Invalid password",
		})
		return
	}

	// create jwt token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": admin.ID,
		"role": admin.Role,
		"email": admin.Email,
		"exp" : time.Now().Add(5*time.Hour).Unix(),
	}).SignedString(jwtKey)

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "Token generation failed",
		})
		return
	}

	c.JSON(200, gin.H{
			"msg": "logged in bro", "token" : token})
	
}