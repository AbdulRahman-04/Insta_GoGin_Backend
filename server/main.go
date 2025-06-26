package main

import (
	"github.com/AbdulRahman-04/Go_Backend_Practice/utils"
	"github.com/gin-gonic/gin"
)


func main(){

	// database import 
    utils.DbConnect()

	router := gin.Default()

	router.GET("/", func (c *gin.Context){
		c.JSON(200, gin.H{
			"msg" :"Hello World in gin",
		})
	})

	

	router.Run(":7070")


}