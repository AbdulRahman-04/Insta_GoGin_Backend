package private

import (
	"context"
	"time"

	"github.com/AbdulRahman-04/Go_Backend_Practice/models"
	"github.com/AbdulRahman-04/Go_Backend_Practice/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var storiesCollection *mongo.Collection

func StoryCollect() {
	storiesCollection = utils.MongoClient.Database("Insta_Backend").Collection("stories")
}

// create story api 
func CreateStory(c*gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId := c.MustGet("userId").(primitive.ObjectID)
	caption := c.PostForm("caption")
	text := c.PostForm("text")
	song := c.PostForm("song")

	imageUrl, err := utils.UploadFile(c)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "file upload failed",
		})
		return
	}

	// create var nd push into db 
	var newStory models.Stories
	newStory.ID = primitive.NewObjectID()
	newStory.UserId = userId
	newStory.Caption = caption
	newStory.Text = text
	newStory.Song = song
	newStory.ImageUrl = imageUrl
	newStory.CreatedAt = time.Now()
	newStory.UpdatedAt = time.Now()

	_, err = storiesCollection.InsertOne(ctx, newStory)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "db error",
		})
		return
	}

	c.JSON(200, gin.H{
			"msg": "story created✅🔥"})
}

// func getall userspecific story 
func GetAllStories(c*gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId := c.MustGet("userId").(primitive.ObjectID)

	cursor, err := storiesCollection.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "db error, no stories found",
		})
		return
	}

	defer cursor.Close(ctx)

	var stories []models.Stories
	err = cursor.All(ctx, &stories)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "error while decoding",
		})
		return
	}

	c.JSON(200, gin.H{
			"msg": "all stories are here👀","stories": stories})
		
}

// get one user specific api 
func GetOneStory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	paramId := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid id",
		})
		return
	}

	var oneStory models.Stories
	err = storiesCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&oneStory)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "no user story found",
		})
		return
	}

	// ✅ Get userId from context (JWT claims)
	userId := c.MustGet("userId").(primitive.ObjectID)

	// ✅ Final response with userId
	c.JSON(200, gin.H{
		"msg":     "user one story is here👀",
		"stories": oneStory,
		"userId":  userId.Hex(),
	})
}


// edit story api 
func EditStory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 🔐 Get user ID from JWT context
	userId := c.MustGet("userId").(primitive.ObjectID)

	// 🆔 Get story ID from URL param
	paramId := c.Param("id")
	storyId, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(400, gin.H{"msg": "Invalid story ID"})
		return
	}

	// 📥 Get form fields
	caption := c.PostForm("caption")
	text := c.PostForm("text")
	song := c.PostForm("song")

	// 📸 Try image upload (optional)
	imageUrl, err := utils.UploadFile(c)
	if err != nil {
		imageUrl = ""
	}

	// 🔍 Check if story exists and belongs to user
	var existingStory models.Stories
	err = storiesCollection.FindOne(ctx, bson.M{
		"_id":     storyId,
		"user_id": userId,
	}).Decode(&existingStory)
	if err != nil {
		c.JSON(404, gin.H{"msg": "Story not found or unauthorized"})
		return
	}

	// 🛠️ Build update object (with $set)
	update := bson.M{
		"$set": bson.M{
			"caption":    caption,
			"text":       text,
			"song":       song,
			"updated_at": time.Now(),
		},
	}

	if imageUrl != "" {
		update["$set"].(bson.M)["imageUrl"] = imageUrl
	}

	// 💾 Update into MongoDB
	_, err = storiesCollection.UpdateOne(ctx, bson.M{"_id": storyId}, update)
	if err != nil {
		c.JSON(500, gin.H{"msg": "DB error while updating story"})
		return
	}

	// ✅ Success
	c.JSON(200, gin.H{"msg": "Story updated successfully ✅"})
}

// delete one story 
func DeleteOneStory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId := c.MustGet("userId").(primitive.ObjectID)
	paramId := c.Param("id")

	storyId, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(400, gin.H{"msg": "Invalid story ID"})
		return
	}

	res, err := storiesCollection.DeleteOne(ctx, bson.M{"_id": storyId, "user_id": userId})
	if err != nil || res.DeletedCount == 0 {
		c.JSON(404, gin.H{"msg": "Story not found or unauthorized"})
		return
	}

	c.JSON(200, gin.H{"msg": "Story deleted successfully ✅"})
}

// delete all stories api
func DeleteAllStories(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId := c.MustGet("userId").(primitive.ObjectID)

	res, err := storiesCollection.DeleteMany(ctx, bson.M{"user_id": userId})
	if err != nil || res.DeletedCount == 0 {
		c.JSON(400, gin.H{"msg": "No stories found or DB error"})
		return
	}

	c.JSON(200, gin.H{"msg": "All your stories deleted ✅"})
}
