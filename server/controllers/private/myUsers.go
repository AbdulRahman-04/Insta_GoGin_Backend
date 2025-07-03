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


var userCollection *mongo.Collection

func UserCollect() {
	userCollection = utils.MongoClient.Database("Insta_Backend").Collection("user")
}


// 🧑‍💼 GET ALL USERS (only admin)
func GetAllUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"msg": "DB error"})
		return
	}

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		c.JSON(500, gin.H{"msg": "Decode error"})
		return
	}

	c.JSON(200, gin.H{"msg": "All users", "users": users})
}

// 👤 GET ONE USER (only admin)
func GetOneUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idParam := c.Param("id")
	userID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(400, gin.H{"msg": "Invalid user ID"})
		return
	}

	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		c.JSON(404, gin.H{"msg": "User not found"})
		return
	}

	c.JSON(200, gin.H{"msg": "User found", "user": user})
}
