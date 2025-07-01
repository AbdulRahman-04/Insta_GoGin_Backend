package private

import (
	"net/http"
	"time"

	"github.com/AbdulRahman-04/Go_Backend_Practice/models"
	"github.com/AbdulRahman-04/Go_Backend_Practice/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost(c *gin.Context) {
	// 🧠 Get userId from JWT context
	userID, _ := c.Get("userId")
	uid := userID.(primitive.ObjectID)

	// ✅ Get caption and tags from form-data
	caption := c.PostForm("caption")
	tags := c.PostFormArray("tags") // multiple tags

	// 📤 Upload file and get path
	imagePath, err := utils.UploadFile(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "file upload failed", "err": err.Error()})
		return
	}

	// 🎯 Create post object
	var newPost models.Post

	newPost.ID = primitive.NewObjectID() // New MongoDB ObjectID
	newPost.UserId = uid                 // User ID from JWT context
	newPost.Caption = caption            // Caption from form-data
	newPost.ImageUrl = imagePath         // Image path from UploadFile
	newPost.Tags = tags                  // Tags from form-data array
	newPost.Likes = 0                    // Default 0 likes
	newPost.CreatedAt = time.Now()       // Current time
	newPost.UpdatedAt = time.Now()       // Same as created for now

	// 💾 Save post in DB
	postCol := utils.MongoClient.Database("Go_Backend_Practice").Collection("posts")
	_, err = postCol.InsertOne(c, newPost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "cannot save post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "post created", "data": newPost})
}
