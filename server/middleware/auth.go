package middleware

import (
	"strings"

	"github.com/AbdulRahman-04/Go_Backend_Practice/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var myJwtKey = []byte(config.AppConfig.JWTKEY)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		// get token in authheader
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(400, gin.H{
				"msg": "no token provided",
			})
			c.Abort()
			return
		}

		// split krdo token ku 
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(400, gin.H{
				"msg": "invalid token format",
			})
			c.Abort()
			return
		}

		// token ku ab store krke verify kro 
		myToken := parts[1]

		token, err := jwt.Parse(myToken, func(token *jwt.Token)(interface{}, error){
			return myJwtKey , nil
		})
		if err != nil {
			c.JSON(400, gin.H{
				"msg": "invalid token",
			})
			c.Abort()
			return
		}

		// claims m poora data store krlo 
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(400, gin.H{
				"msg": "no data found",
			})
			c.Abort()
			return
		}

		// user/admin id nikalo aur mongoid m convert kro 
		userStrId, ok := claims["id"].(string)
		if !ok {
			c.JSON(400, gin.H{
				"msg": "no id found",
			})
			c.Abort()
			return
		}

		userId, err := primitive.ObjectIDFromHex(userStrId)
		if err != nil {
			c.JSON(400, gin.H{
				"msg": "couldn't parse id into mongo id",
			})
			c.Abort()
			return
		}

		// role nikalo 
		role, ok := claims["role"].(string)
		if !ok {
			c.JSON(400, gin.H{
				"msg": "no role found",
			})
			c.Abort()
			return
		} 

		// context m set krdo role nd id
		c.Set("userId", userId)
		c.Set("role", role)

		c.Next()
	}
}