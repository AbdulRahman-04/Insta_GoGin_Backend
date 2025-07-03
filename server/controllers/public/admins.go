package public

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
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

var adminCollection *mongo.Collection

var jwtKey = []byte(config.AppConfig.JWTKEY)

var url = config.AppConfig.URL

func AdminCollect() {
	adminCollection = utils.MongoClient.Database("GO_BACKEND_practice").Collection("admin")
}

func adminTokenGenerate(length int) string {
	d := make([]byte, length)
	_, _ = rand.Read(d)
	return hex.EncodeToString(d)
}

// admin signup api
func AdminSignUp(c *gin.Context) {
	// create ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// type struct
	type AdminSignUp struct {
		AdminName string `json:"adminname" form:"adminname"`
		Email     string `json:"email" form:"email"`
		Password  string `json:"password" form:"password"`
		Age       int    `json:"age" form:"age"`
		Phone     string `json:"phone" form:"phone"`
		Language  string `json:"language" form:"language"`
	}

	// create var nd bint it
	var inputAdmin AdminSignUp

	if err := c.ShouldBindJSON(&inputAdmin); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
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

	if !strings.Contains(inputAdmin.Email, "@") {
		c.JSON(400, gin.H{
			"msg": "invalid email",
		})
		return
	}

	if len(inputAdmin.Password) < 6 {
		c.JSON(400, gin.H{
			"msg": "invalid pass length",
		})
		return
	}

	if len(inputAdmin.Phone) < 10 {
		c.JSON(400, gin.H{
			"msg": "invalid number length",
		})
		return
	}

	// duplicate check
	count, err := adminCollection.CountDocuments(ctx, bson.M{"email": inputAdmin.Email})
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid db error",
		})
		return
	}

	if count > 0 {
		c.JSON(400, gin.H{
			"msg": "admin already exists",
		})
		return
	}

	// hash the pass
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(inputAdmin.Password), 10)
	emailToken := adminTokenGenerate(8)
	phoneToken := adminTokenGenerate(8)

	// create admin model var and push values inside it
	var admin models.Admin

	admin.AdminName = inputAdmin.AdminName
	admin.Email = inputAdmin.Email
	admin.Password = string(hashPass)
	admin.Role = "admin"
	admin.Phone = inputAdmin.Phone
	admin.Age = inputAdmin.Age
	admin.Language = inputAdmin.Language
	admin.AdminVerified.Email = false
	admin.AdminVerfiyToken.Email = emailToken
	admin.AdminVerfiyToken.Phone = phoneToken
	admin.CreatedAt = time.Now()
	admin.UpdatedAt = time.Now()

	// send email link
	go func() {

		emailData := utils.EmailData{
			From:    "team insta",
			To:      inputAdmin.Email,
			Subject: "email verify",
			Html:    fmt.Sprintf(`<a href="%s/api/public/admin/emailverify/%s">Click here to verify your email</a>`, URL, emailToken),
		}
		_ = utils.SendEmail(emailData)
	}()

	// insert into db
	_, err = adminCollection.InsertOne(ctx, admin)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldn't add data",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "Admin registered, please verify ur email nd login",
	})

}

// email verify api
func EmailAdminVerify(c *gin.Context) {
	// ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// take token from param
	token := c.Param("token")

	// compare the param token with each document in db
	var admin models.Admin
	err := adminCollection.FindOne(ctx, bson.M{
		"adminVerifyToken.emailToken": token,
	}).Decode(&admin)

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid token",
		})
		return
	}

	// check if the link hasnt been clicked more than once
	if admin.AdminVerified.Email {
		c.JSON(400, gin.H{
			"msg": "admin email already verified",
		})
		return
	}

	//  update the db
	update := bson.M{
		"$set": bson.M{
			"adminVerified.emailVerified": true,
			"adminVerifyToken.emailToken": nil,
			"updated_at":                  time.Now(),
		}}

	// update the changes into db
	_, err = adminCollection.UpdateByID(ctx, admin.ID, update)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldn't update email",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "admin email verified",
	})
}

// signin api admin
func AdminSignIn(c *gin.Context) {
	// ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// type struct
	type AdminSignIn struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	// create a var and bind into it
	var inputAdmin AdminSignIn

	if err := c.ShouldBindJSON(&inputAdmin); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validations
	if inputAdmin.Email == "" || inputAdmin.Password == "" {
		c.JSON(400, gin.H{
			"msg": "fill all fields",
		})
		return
	}

	if !strings.Contains(inputAdmin.Email, "@") {
		c.JSON(400, gin.H{
			"msg": "invalid email",
		})
		return
	}

	if len(inputAdmin.Password) < 6 {
		c.JSON(400, gin.H{
			"msg": "invalid pass",
		})
		return
	}

	// find admin in db using admin model
	var admin models.Admin

	err := adminCollection.FindOne(ctx, bson.M{"email": inputAdmin.Email}).Decode(&admin)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldn't find email pls register",
		})
		return
	}

	// compare pass
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(inputAdmin.Password)); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid pass",
		})
		return
	}

	// token generation
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    admin.ID,
		"role":  admin.Role,
		"email": admin.Email,
		"exp":   time.Now().Add(5 * time.Hour).Unix(),
	}).SignedString(jwtKey)

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldn't generate token",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "user logged in successfully!✅", "token": token})
}

// change password
func ChangeAdminPass(c *gin.Context) {
	// ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// type changepass struct
	type ChangeAdminPass struct {
		Email       string `json:"email" form:"email"`
		OldPassword string `json:"oldpassword" form:"oldpassword"`
		NewPassword string `json:"newpassword" form:"newpassword"`
	}

	// create a var and bind into it
	var inputAdmin ChangeAdminPass

	if err := c.ShouldBindJSON(&inputAdmin); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validations
	if inputAdmin.Email == "" || inputAdmin.OldPassword == "" || inputAdmin.NewPassword == "" {
		c.JSON(400, gin.H{
			"msg": "invalid input, pls fill all",
		})
		return
	}

	if !strings.Contains(inputAdmin.Email, "@") {
		c.JSON(400, gin.H{
			"msg": "invalid email",
		})
		return
	}

	if len(inputAdmin.NewPassword) < 6 {
		c.JSON(400, gin.H{
			"msg": "invalid new pass length",
		})
		return
	}

	// find the email in db
	var admin models.Admin

	err := adminCollection.FindOne(ctx, bson.M{"email": inputAdmin.Email}).Decode(&admin)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "not found email",
		})
		return
	}

	// compare old pass
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(inputAdmin.OldPassword))
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid pass",
		})
		return
	}

	// hash the new pass
	hashNewPass, err := bcrypt.GenerateFromPassword([]byte(inputAdmin.NewPassword), 10)

	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldnt hash pass",
		})
		return
	}

	// update into db
	update := bson.M{
		"$set": bson.M{
			"password":   string(hashNewPass),
			"updated_at": time.Now(),
		}}

	// add the changes into db
	_, err = adminCollection.UpdateByID(ctx, admin.ID, update)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "db error",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "Password changes successfully!✅",
	})

}

// forgot password api
func ForgotPass(c *gin.Context) {
	// ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// type struct
	type ForgotPass struct {
		Email string `json:"email" form:"email"`
	}

	// create var and bind it
	var inputAdmin ForgotPass

	if err := c.ShouldBindJSON(&inputAdmin); err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// validation
	if !strings.Contains(inputAdmin.Email, "@") || inputAdmin.Email == "" {
		c.JSON(400, gin.H{
			"msg": "either mail is incorrect of empty",
		})
		return
	}

	// check admin email in db
	var admin models.Admin

	err := adminCollection.FindOne(ctx, bson.M{"email": inputAdmin.Email}).Decode(&admin)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid! no admin found",
		})
		return
	}

	// generate a new pass
	tempPass := adminTokenGenerate(8)
	hashTempPass, err := bcrypt.GenerateFromPassword([]byte(tempPass), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "couldnt hash pass",
		})
		return
	}

	update := bson.M{"$set": bson.M{"password": string(hashTempPass), "updated_at": time.Now()}}
	_, err = adminCollection.UpdateByID(ctx, admin.ID, update)
	if err != nil {
		c.JSON(400, gin.H{"msg": "Password update failed!"})
		return
	}

	// Send temp password to email
	emailData := utils.EmailData{
		From:    "Team WebXpertz 💌",
		To:      inputAdmin.Email,
		Subject: "🔑 Temporary Password",
		Html:    fmt.Sprintf(`<h2>Password Reset</h2><p>Dear %s, your temporary password is: <strong>%s</strong></p><p>Please log in and change it immediately.</p>`, admin.AdminName, tempPass),
	}
	_ = utils.SendEmail(emailData)

	c.JSON(200, gin.H{"msg": "Temporary password sent to email ✅"})
}
