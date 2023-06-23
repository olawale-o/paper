package articles

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Article struct {
	TITLE  string `json:"title"`
	AUTHOR string `json:"author"`
}

var a []Article = []Article{
	{TITLE: "Article 1", AUTHOR: "Author 1"},
	{TITLE: "Article 2", AUTHOR: "Author 2"},
	{TITLE: "Article 3", AUTHOR: "Author 3"},
	{TITLE: "Article 4", AUTHOR: "Author 4"},
}

func GetArticles(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	//  "articles": articles,
	// })
	c.IndentedJSON(http.StatusOK, a)
}

func NewArticle(c *gin.Context) {
	var newArticle Article
	if err := c.BindJSON(&newArticle); err != nil {
		log.Println(err)
		return
	}
	a = append(a, newArticle)
	c.IndentedJSON(http.StatusCreated, newArticle)
}

func ShowArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, article := range a {
		if i == id {
			c.IndentedJSON(http.StatusOK, article)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
}

func UpdateArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedArticle Article
	if err := c.BindJSON(&updatedArticle); err != nil {
		log.Println(err)
		return
	}
	for i := range a {
		if i == id {
			a[i] = updatedArticle
			c.IndentedJSON(http.StatusOK, updatedArticle)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range a {
		if i == id {
			a = append(a[:i], a[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Article deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
}
