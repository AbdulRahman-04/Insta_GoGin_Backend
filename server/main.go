package main

import (
	"fmt"

	"github.com/AbdulRahman-04/Go_Backend_Practice/config"
	"github.com/AbdulRahman-04/Go_Backend_Practice/utils"
	"github.com/gin-gonic/gin"
)


func main() {
	// DBCONNECT 
	utils.DbConnect()

	router := gin.Default()

	router.GET("/", func (c *gin.Context){
		c.JSON(200, gin.H{
			"msg" : "Hello From Gin",
		})
	})

	router.Run(fmt.Sprintf(":%d", config.AppConfig.Port))
}