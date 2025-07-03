package routes

import (
	"github.com/AbdulRahman-04/Go_Backend_Practice/controllers/private"
	"github.com/AbdulRahman-04/Go_Backend_Practice/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPrivateRoutes(router *gin.Engine) {
	privateGroup := router.Group("/api/private")

	// 🔐 JWT Middleware for all private routes
	privateGroup.Use(middleware.AuthMiddleware())

	// ✅ Post routes
	privateGroup.POST("/addpost", middleware.OnlyUser(), private.CreatePost)
	privateGroup.GET("/getallposts", middleware.OnlyUser(), private.GetAllPosts)
	privateGroup.GET("/getonepost/:id", middleware.OnlyUser(), private.GetOnePost)
	privateGroup.PUT("/editpost/:id", middleware.OnlyUser(), private.EditPost)
	privateGroup.DELETE("/deletepost/:id", middleware.OnlyUser(), private.DeleteOnePost)
	privateGroup.DELETE("/deleteallposts", middleware.OnlyUser(), private.DeleteAllPosts)

	// ✅ Story routes
	privateGroup.POST("/addstory", middleware.OnlyUser(), private.CreateStory)
	privateGroup.GET("/getallstories", middleware.OnlyUser(), private.GetAllStories)
	privateGroup.GET("/getonestories/:id", middleware.OnlyUser(), private.GetOneStory)
	privateGroup.PUT("/editonestories/:id", middleware.OnlyUser(), private.EditStory)
	privateGroup.DELETE("/deleteonestories/:id", middleware.OnlyUser(), private.DeleteOneStory)
	privateGroup.DELETE("/deleteallstories", middleware.OnlyUser(), private.DeleteAllStories)

	// ✅ Admin-only user routes (RBAC check)
	privateGroup.GET("/users", middleware.OnlyAdmin(), private.GetAllUsers)
	privateGroup.GET("/users/:id", middleware.OnlyAdmin(), private.GetOneUser)
}
