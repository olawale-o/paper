package controller

import (
	"comments/model"
	"comments/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func New(c *gin.Context) {
	articleId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Please provide valid credntials"})
		return
	}
	err, _ := service.NewComment(comment, articleId)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Comment saved"})
}

func Show(c *gin.Context) {
	articleId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	commentId, _ := primitive.ObjectIDFromHex(c.Param("cid"))

	err, res := service.GetComment(articleId, commentId)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "comment", "data": res})
}

func Index(c *gin.Context) {
	articleId, _ := primitive.ObjectIDFromHex(c.Param("id"))

	res := service.GetComments(articleId)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "comments", "data": res})
}
