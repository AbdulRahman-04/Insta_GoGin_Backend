package middleware

import (
	"strings"

	"github.com/AbdulRahman-04/Go_Backend_Practice/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// jwt key lalo config se
var myKey = []byte(config.AppConfig.JWTKEY)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the token in auth header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(400, gin.H{
				"msg": "please provide a token",
			})
			c.Abort()
			return
		}

		// token ku split kro
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(400, gin.H{
				"msg": "invalid token format",
			})
			c.Abort()
			return
		}

		// ab ye token ku store krlo new var m
		myToken := parts[1]

		// ab token ku parse krke verify kro
		token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
			return myKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(400, gin.H{
				"msg": "invalid token",
			})
			c.Abort()
			return
		}

		// ab jo h so token k andar ka data nikalo claims var m
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(400, gin.H{
				"msg": "no data found",
			})
			c.Abort()
			return
		}

		// ab jo h userid/admin id extract krke mongodb id m convert kro
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
				"msg": "invalid id format",
			})
			c.Abort()
			return
		}

		// role nikalo claims m se
		role, ok := claims["role"].(string)
		if !ok {
			c.JSON(400, gin.H{
				"msg": "no role found",
			})
			c.Abort()
			return
		}

		// set krdo context m
		c.Set("userId", userId)
		c.Set("role", role)

		c.Next()
	}
}
