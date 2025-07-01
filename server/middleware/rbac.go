package middleware

import "github.com/gin-gonic/gin"

func OnlyAdmin() gin.HandlerFunc{
	return func (c *gin.Context){
		role, exists := c.Get("role")
		if !exists || role != "admin"{
			c.JSON(400, gin.H{
				"msg": "admin access only",
			})
          c.Abort()
		  return
		}
		c.Next()
	}
}

func OnlyUser() gin.HandlerFunc{
	return func (c *gin.Context){
		role, exists := c.Get("role")
		if !exists || role!= "user"{
			c.JSON(400, gin.H{
				"msg": "users access only",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}