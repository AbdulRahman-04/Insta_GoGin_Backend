// package private

// import (
// 	"context"
// 	"net/http"
// 	"time"

// 	"github.com/AbdulRahman-04/Go_Backend_Practice/models"
// 	"github.com/AbdulRahman-04/Go_Backend_Practice/utils"
// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// // 🗂️ Collection global
// var postCollection *mongo.Collection

// func PostCollect() {
// 	postCollection = utils.MongoClient.Database("GO_BACKEND_practice").Collection("posts")
// }

// //////////////////////////////////////
// // 📌 1. CREATE POST
// func CreatePost(c *gin.Context) {
// 	uid := c.MustGet("userId").(primitive.ObjectID)

// 	caption := c.PostForm("caption")
// 	tags := c.PostFormArray("tags")

// 	imagePath, err := utils.UploadFile(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"msg": "image upload failed"})
// 		return
// 	}

// 	var newPost models.Post
// newPost.ID = primitive.NewObjectID()
// newPost.UserId = uid
// newPost.Caption = caption
// newPost.ImageUrl = imagePath
// newPost.Tags = tags
// newPost.Likes = 0
// newPost.CreatedAt = time.Now()
// newPost.UpdatedAt = time.Now()

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	_, err = postCollection.InsertOne(ctx, newPost)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"msg": "couldn't create post"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"msg": "post created", "data": newPost})
// }

// //////////////////////////////////////
// // 📌 2. GET ALL POSTS (User specific)
// func GetAllPosts(c *gin.Context) {
// 	uid := c.MustGet("userId").(primitive.ObjectID)

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	cursor, err := postCollection.Find(ctx, bson.M{"user_id": uid})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"msg": "error getting posts"})
// 		return
// 	}
// 	defer cursor.Close(ctx)

// 	var posts []models.Post
// 	if err = cursor.All(ctx, &posts); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"msg": "error decoding posts"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": posts})
// }

// //////////////////////////////////////
// // 📌 3. GET ONE POST BY ID
// func GetOnePost(c *gin.Context) {
// 	postID := c.Param("id")

// 	objID, err := primitive.ObjectIDFromHex(postID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid post id"})
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	var post models.Post
// 	err = postCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&post)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"msg": "post not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": post})
// }

// //////////////////////////////////////
// // 📌 4. EDIT ONE POST
// func EditPost(c *gin.Context) {
// 	postID := c.Param("id")
// 	objID, err := primitive.ObjectIDFromHex(postID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid post id"})
// 		return
// 	}

// 	caption := c.PostForm("caption")
// 	tags := c.PostFormArray("tags")

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	update := bson.M{
// 		"$set": bson.M{
// 			"caption":    caption,
// 			"tags":       tags,
// 			"updated_at": time.Now(),
// 		},
// 	}

// 	_, err = postCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"msg": "failed to update post"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"msg": "post updated"})
// }

// //////////////////////////////////////
// // 📌 5. DELETE ONE POST
// func DeleteOnePost(c *gin.Context) {
// 	postID := c.Param("id")
// 	objID, err := primitive.ObjectIDFromHex(postID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid post id"})
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	_, err = postCollection.DeleteOne(ctx, bson.M{"_id": objID})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"msg": "failed to delete post"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"msg": "post deleted"})
// }

// //////////////////////////////////////
// // 📌 6. DELETE ALL POSTS BY USER
// func DeleteAllPosts(c *gin.Context) {
// 	uid := c.MustGet("userId").(primitive.ObjectID)

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	_, err := postCollection.DeleteMany(ctx, bson.M{"user_id": uid})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"msg": "failed to delete all posts"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"msg": "all posts deleted"})
// }

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
	imageUrl, err := utils.UploadFile(c)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "error uploading file",
		})
		return
	}

	// var m daalke db m add krdo 
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
			"msg": "new post created",
		})
}

func GetAllPosts(c*gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// get user id 
	userId := c.MustGet("userId").(primitive.ObjectID)

	// cursor connection bnao db se connect hoke find krne 
	cursor , err := postCollection.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "error finding user posts in db",
		})
		return
	}

	defer cursor.Close(ctx)

	var posts []models.Post
	err = cursor.All(ctx, &posts)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "error decoding",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "all posts are here", "Posts": posts})
}

// get one by id api 
func GetOnePost(c*gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	paramId := c.Param("id")

	mongoId, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid post id",
		})
		return
	}

	var onePost models.Post
	err = postCollection.FindOne(ctx, bson.M{"_id": mongoId}).Decode(&onePost)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "errordb",
		})
		return
	}

	c.JSON(200, gin.H{
			"msg": "user one post is here","post": onePost})
		return
}