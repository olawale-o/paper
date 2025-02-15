package controller

import (
	"comments/events"
	"comments/model"
	"comments/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func Event(c *gin.Context) {
	log.Println("Consuming comments")
	var comment model.Payload
	if err := c.BindJSON(&comment); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Please provide valid credntials"})
		return
	}
	events.ConsumeEvent(comment)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Comment saved"})
}
