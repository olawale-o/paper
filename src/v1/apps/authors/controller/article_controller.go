package controller

import (
	"authors/model"
	"authors/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var articleCollection = client.Database("go").Collection("articles")

func ArticleIndex(c *gin.Context) {
	authorId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	fmt.Println(authorId)
	res, _ := service.AllArticles(authorId)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article retrieved successfully", "article": res})
}

func ArticleUpdate(c *gin.Context) {
	authorId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	articleId, _ := primitive.ObjectIDFromHex(c.Param("articleId"))

	var article model.Article
	if err := c.BindJSON(&article); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Unable to process entities"})
		return
	}

	res, err := service.UpdateArticle(article, authorId, articleId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article updated successfully", "article": res})
}

func ArticleDelete(c *gin.Context) {
	authorId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	articleId, _ := primitive.ObjectIDFromHex(c.Param("articleId"))

	res, err := service.DeleteArticle(authorId, articleId)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article deleted successfully", "article": res})
}
