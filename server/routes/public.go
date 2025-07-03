package routes

import (
	"github.com/AbdulRahman-04/Go_Backend_Practice/controllers/public"
	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(router *gin.Engine) {
	publicGroup := router.Group("/api/public")

	// 👤 USER PUBLIC ROUTES
	userGroup := publicGroup.Group("/user")
	userGroup.POST("/register", public.UserSignUp)
	userGroup.POST("/login", public.UserSignIn)
	userGroup.GET("/emailverify/:token", public.UserEmailVerify)
	userGroup.POST("/change-password", public.ChangeUserPass)
	userGroup.POST("/forgot-password", public.ForgotPassUser)

	// 🛡️ ADMIN PUBLIC ROUTES
	adminGroup := publicGroup.Group("/admin")
	adminGroup.POST("/register", public.AdminSignUp)
	adminGroup.POST("/login", public.AdminSignIn)
	adminGroup.GET("/emailverify/:token", public.EmailAdminVerify)
	adminGroup.POST("/change-password", public.ChangeAdminPass)
	adminGroup.POST("/forgot-password", public.ForgotPass)

	// ✅ Optional Ping Route
	publicGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "✅ API is live!"})
	})
}
