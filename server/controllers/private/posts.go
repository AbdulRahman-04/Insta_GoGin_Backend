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

var postCollection *mongo.Collection

func PostCollect(){
	postCollection = utils.MongoClient.Database("GO_BACKEND_practice").Collection("post")
}

func CreatePost(c*gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId := c.MustGet("userId").(primitive.ObjectID)
	caption := c.PostForm("caption")
	tags := c.PostFormArray("tags")
	imageUrl , err := utils.UploadFile(c)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "error while uploading file",
		})
		return
	}
	// var bnake db m push krdo 
	var newPost models.Post

	newPost.ID = primitive.NewObjectID()
	newPost.UserId = userId
	newPost.Caption = caption
	newPost.Tags = tags
	newPost.ImageUrl = imageUrl
	newPost.CreatedAt = time.Now()
	newPost.UpdatedAt = time.Now()

	_, err = postCollection.InsertOne(ctx, newPost)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "errordb",
		})
		return
	}

	c.JSON(200, gin.H{
			"msg": "post created✅",
		})
}

func GetAllPosts(c*gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId := c.MustGet("userId").(primitive.ObjectID)

	cursor, err := postCollection.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "errordb",
		})
		return
	}

	// decode krke array m store krke show kro repsonse m 
	var posts []models.Post
	err = cursor.All(ctx, &posts)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "error while decoding",
		})
		return
	}

	c.JSON(200, gin.H{
			"msg": "all posts are here", "posts": posts})
		return
}

// normal get all posts 
func GetAllUserPosts(c*gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := postCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "errordb",
		})
		return
	}

	var allPosts []models.Post
	err = cursor.All(ctx, &allPosts)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "errordb",
		})
		return
	}

	c.JSON(200, gin.H{
			"msg": "all users posts are here",
		})
}

// get one post 
func GetOnePost(c*gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	paramId := c.Param("id")

	mongoId, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "errordb",
		})
		return
	}

	// find the id in db 
	var onePost models.Post
	err = postCollection.FindOne(ctx, bson.M{"_id":mongoId}).Decode(&onePost)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "errordb",
		})
		return
	}

	c.JSON(200, gin.H{
			"msg": "one post of user is here", "post": onePost})
		
}


// get one public api 
func GetOnePostUser(c*gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	paramId := c.Param("id")

	mongoId, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid id",
		})
		return
	}

	var onePost models.Post
	err = postCollection.FindOne(ctx, bson.M{"_id": mongoId}).Decode(&onePost)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid id, not found in db",
		})
		return
	}
}


// EDIT USER POST API

func EditPost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 🔐 Get user ID from context (JWT)
	userId := c.MustGet("userId").(primitive.ObjectID)

	// 🆔 Get Post ID from URL param
	paramId := c.Param("id")
	postId, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "Invalid post ID format",
		})
		return
	}

	// 📥 Get form data
	caption := c.PostForm("caption")
	tags := c.PostFormArray("tags")

	// 📸 Try to upload image (optional)
	imageUrl, err := utils.UploadFile(c)
	if err != nil {
		imageUrl = "" // optional image
	}

	// 🔍 Check if post exists & belongs to current user
	var existingPost models.Post
	err = postCollection.FindOne(ctx, bson.M{
		"_id":     postId,
		"user_id": userId,
	}).Decode(&existingPost)
	if err != nil {
		c.JSON(404, gin.H{
			"msg": "Post not found or unauthorized",
		})
		return
	}

	// 🔧 Prepare update object with `$set` (Method 1 style)
	update := bson.M{
		"$set": bson.M{
			"caption":    caption,
			"tags":       tags,
			"updated_at": time.Now(),
		},
	}

	if imageUrl != "" {
		update["$set"].(bson.M)["image_url"] = imageUrl
	}

	// 💾 Update in MongoDB
	_, err = postCollection.UpdateOne(ctx, bson.M{"_id": postId}, update)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "DB error while updating post",
		})
		return
	}

	// ✅ Success
	c.JSON(200, gin.H{
		"msg": "Post updated successfully ✅",
	})
}

// delete one post 
func DeleteOnePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId := c.MustGet("userId").(primitive.ObjectID)
	paramId := c.Param("id")

	postId, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(400, gin.H{"msg": "Invalid post ID"})
		return
	}

	// 🔐 Ensure post belongs to this user
	res, err := postCollection.DeleteOne(ctx, bson.M{"_id": postId, "user_id": userId})
	if err != nil || res.DeletedCount == 0 {
		c.JSON(404, gin.H{"msg": "Post not found or unauthorized"})
		return
	}

	c.JSON(200, gin.H{"msg": "Post deleted successfully ✅"})
}

// delete all posts 
func DeleteAllPosts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId := c.MustGet("userId").(primitive.ObjectID)

	res, err := postCollection.DeleteMany(ctx, bson.M{"user_id": userId})
	if err != nil || res.DeletedCount == 0 {
		c.JSON(400, gin.H{"msg": "No posts found or DB error"})
		return
	}

	c.JSON(200, gin.H{"msg": "All your posts deleted ✅"})
}

