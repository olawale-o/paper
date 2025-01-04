package middlewares

import (
	"fmt"
	"go-simple-rest/src/v1/jwt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")

		if err != nil {
			fmt.Println("Token is missing")
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Missing token"})
			return
		}

		token, err := jwt.VerifyToken(tokenString)
		if err != nil {
			fmt.Printf("Token verification failed: %v\\n", err)
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Unathorized"})
			return
		}
		fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)
		// before request

		c.Next()

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)

	}
}
