package main

import (
	"fmt"

	"github.com/AbdulRahman-04/Go_Backend_Practice/config"
	"github.com/AbdulRahman-04/Go_Backend_Practice/controllers/private"
	"github.com/AbdulRahman-04/Go_Backend_Practice/controllers/public"
	"github.com/AbdulRahman-04/Go_Backend_Practice/routes" // 👈 import routes
	"github.com/AbdulRahman-04/Go_Backend_Practice/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// 🛢️ DB CONNECT
	utils.DbConnect()

	// 🚀 INIT GIN
	router := gin.Default()

	// collects functions call
	private.UserCollect()
	public.UserCollect()
	public.AdminCollect()
	private.PostCollect()
    private.StoryCollect()

	// ✅ ROOT CHECK
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "HELLO FROM GIN"})
	})

	// 📡 REGISTER ROUTES
	routes.RegisterPublicRoutes(router)   // 👈 for /api/public/*
	routes.RegisterPrivateRoutes(router)  // 👈 for /api/private/*

	// 🔥 RUN SERVER
	router.Run(fmt.Sprintf(":%d", config.AppConfig.Port))
}
