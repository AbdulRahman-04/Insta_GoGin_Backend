package middleware

import (
	"github.com/gin-gonic/gin"
    "fmt"
  )


func OnlyAdmin() gin.HandlerFunc{
	return func(c*gin.Context){
		role, exists := c.Get("role")
		if !exists || role != "admin"{
			c.JSON(400, gin.H{
				"msg": "Access denied, only admins allowed here",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func OnlyUser() gin.HandlerFunc {
	return func(c*gin.Context){
		role, exists := c.Get("role")
		if !exists || role != "user"{
			c.JSON(400, gin.H{
				"msg": "Access denied, only users allowed here",
			})
			c.Abort()
			return
		}
		fmt.Println("Middleware hit")

		c.Next()
	}
}