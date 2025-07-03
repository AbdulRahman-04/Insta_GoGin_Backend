package private

import (
	"context"
	"time"

	"github.com/AbdulRahman-04/Go_Backend_Practice/models"
	"github.com/AbdulRahman-04/Go_Backend_Practice/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var storiesCollection *mongo.Collection

func storyCollect() {
	storiesCollection = utils.MongoClient.Database("").Collection("stories")
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
			"msg": "story created"})
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
			"msg": "all stories are here","stories": stories})
		return
}

// normal get all stories 
func GetAllUserStories(c*gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor , err := storiesCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "db error",
		})
		return
	}

	defer cursor.Close(ctx)

	var AllStories []models.Stories
	err = cursor.All(ctx, &AllStories)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "db error decoding",
		})
		return
	}

	c.JSON(200, gin.H{
			"msg": "all stories are here", "stories": AllStories})

}

// get one user specific api 
func GetOneStory(c*gin.Context){
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
	err = storiesCollection.FindOne(ctx, bson.M{"user_id": objId}).Decode(&oneStory)
 
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "no user story found",
		})
		return
	}

	c.JSON(200, gin.H{
			"msg": "user stories are here","stories": oneStory})
}

// normal get one api 
func GetOneStoryUser(c*gin.Context){
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
			"msg": "db error, no such story found",
		})
		return
	}

	c.JSON(200, gin.H{
			"msg": "story is here", "story": oneStory})
		return
}