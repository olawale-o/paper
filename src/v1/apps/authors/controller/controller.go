package controller

import (
	"authors/db"
	"authors/model"
	"authors/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client, ctx, err = db.Connect()

var userCollection = client.Database("go").Collection("users")

func Index(c *gin.Context) {
	res, err := service.ShowAuthors()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Retrieved all authors", "data": res})
}
func Show(c *gin.Context) {
	authorId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	res, err := service.ShowAuthor(authorId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Retrieved single authors", "data": res})
}

func Update(c *gin.Context) {
	authorId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	var author model.Author
	if err := c.BindJSON(&author); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Unable to process entities"})
		return
	}

	res, err := service.UpdateAuthor(authorId, author)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Author updated successfully", "data": res})
}

func Delete(c *gin.Context) {
	authorId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	res, err := service.DeleteAuthor(authorId)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Author deleted successfully", "data": res})
}
