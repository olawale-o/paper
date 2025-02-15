package controller

import (
	"articles/model"
	"articles/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	// userId := c.MustGet("userId").(string)
	// log.Printf("userId %s", userId)
	articles := service.GetAll()
	c.IndentedJSON(http.StatusOK, articles)
}

func NewArticle(c *gin.Context) {
	_, id := service.CreateArticle(c)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Article created", "id": id})
}

func ShowArticle(c *gin.Context) {
	article, err := service.GetArticle(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"article": article})
}

func UpdateArticle(c *gin.Context) {
	result, err := service.Update(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article updated", "article": result})
}

func DeleteArticle(c *gin.Context) {
	result, err := service.Delete(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article deleted", "article": result})
}

// func ArticleComments(c *gin.Context) {
// 	articleId := c.Param("id")
// 	comments := service.GetComments(articleId)
// 	c.IndentedJSON(http.StatusOK, comments)
// }

func NewComment(c *gin.Context) {
	articleId := c.Param("id")

	comment := model.ArticleComment{}
	if err := c.BindJSON(&comment); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Please provide valid data"})
		return
	}
	service.CreateComment(articleId, comment)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Comment created"})
}
