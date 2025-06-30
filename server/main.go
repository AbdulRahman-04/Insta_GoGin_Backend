package main

import "github.com/gin-gonic/gin"


func main(){
	router := gin.Default()

	router.GET("/", func (c *gin.Context)  {
		c.JSON(200, gin.H{
			"msg": "HELLO FROM GIN	",
		})
	})

	router.Run(":6060")
}