package main

import "github.com/gin-gonic/gin"


func main(){
	router := gin.Default()

	router.GET("/", func (c *gin.Context){
		c.JSON(200, gin.H{
			"msg" :"Hello World in gin",
		})
	})

	router.Run(":7070")
}