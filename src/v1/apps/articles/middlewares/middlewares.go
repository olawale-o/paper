package middlewares

import (
	"articles/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")

		if err != nil {
			fmt.Println("Token is missing")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing token"})
			return
		}

		token, err := jwt.VerifyToken(tokenString)
		if err != nil {
			fmt.Printf("Token verification failed: %v\\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unathorized"})
			return
		}
		fmt.Printf("Token verified successfully. Claims: %+v\\n", token)
		// before request
		//
		fmt.Println()

		fmt.Printf("Username %v", token)

		c.Set("userId", token)

		c.Next()

		// access the status we are sending
		// status := c.Writer.Status()
		// log.Println(status)

	}
}
