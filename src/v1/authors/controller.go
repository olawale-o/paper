package authors

import (
	"go-simple-rest/db"
	"go-simple-rest/src/v1/articles"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client, ctx, err = db.Connect()

var articleCollection = client.Database("go").Collection("articles")
var userCollection = client.Database("go").Collection("users")

func CreateArticle(c *gin.Context) {
	authorId, _ := primitive.ObjectIDFromHex(c.MustGet("userId").(string))
	var newArticle articles.Article
	if err := c.BindJSON(&newArticle); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Unable to process entities"})
		return
	}
	res, err := Create(newArticle, authorId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article created successfully", "article": res})
}

func UpdateArticle(c *gin.Context) {
	authorId, _ := primitive.ObjectIDFromHex(c.MustGet("userId").(string))
	articleId, _ := primitive.ObjectIDFromHex(c.Param("articleId"))

	var article articles.Article
	if err := c.BindJSON(&article); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Unable to process entities"})
		return
	}

	res, err := Update(article, authorId, articleId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article updated successfully", "article": res})
}

func DeleteArticle(c *gin.Context) {
	authorId, _ := primitive.ObjectIDFromHex(c.MustGet("userId").(string))
	articleId, _ := primitive.ObjectIDFromHex(c.Param("articleId"))

	res, err := Delete(authorId, articleId)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article deleted successfully", "article": res})
}
