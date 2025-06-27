package main

import (
	"fmt"

	"github.com/AbdulRahman-04/Go_Backend_Practice/config"
	"github.com/AbdulRahman-04/Go_Backend_Practice/utils"
	"github.com/gin-gonic/gin"
)

func main(){
	// DB IMPORT
	utils.DbConnect()
 
	router := gin.Default()

	router.GET("/", func (c *gin.Context){
		c.JSON(200, gin.H{
			"msg" : "Hello from gin",
		})
	})

	router.Run(fmt.Sprintf(":%d", config.AppConfig.Port))
	 

}