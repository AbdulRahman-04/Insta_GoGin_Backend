package main

import (
	"github.com/AbdulRahman-04/Go_Backend_Project/utils"
	"github.com/gin-gonic/gin"
)


func main(){
	// databse import 
	utils.DbConnect()
	
	router := gin.Default()

	router.GET("/", func (c *gin.Context){
		c.JSON(200, gin.H{
			"msg" : "HELLO FROM GIN",
		})
	})

	router.Run(":3050")
}