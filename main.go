package main

import (
	"go-simple-rest/src"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	src.Routes(r)
	log.Println("Starting server... on port ", 8080)
	r.Run("localhost:8080")
}
