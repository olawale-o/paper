package authors

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AuthorArticleIndex godoc
// @Tags Author
// @Summary Get all articles by author
// @Description Retrieves all articles written by a specific author.
// @Param id path string true "Author ID"
// @Produce json
// @Success 200 {object} model.AuthorArticleResponse "Response"
// @Failure 400 {object} string "Error"
// @Failure 500 {object} string "Error"
// @Router /authors/{id}/articles [get]
func ArticleIndex(c *gin.Context) {
	authorId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	fmt.Println(authorId)
	res, _ := AllArticles(authorId)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article retrieved successfully", "articles": res})
}

// func ArticleNew(c *gin.Context) {
// 	authorId, _ := primitive.ObjectIDFromHex(c.Param("id"))
// 	var newArticle articles.Article
// 	if err := c.BindJSON(&newArticle); err != nil {
// 		log.Println(err)
// 		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Unable to process entities"})
// 		return
// 	}
// 	res, err := CreateArtcile(newArticle, authorId)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article created successfully", "article": res})
// }

// func ArticleUpdate(c *gin.Context) {
// 	authorId, _ := primitive.ObjectIDFromHex(c.Param("id"))
// 	articleId, _ := primitive.ObjectIDFromHex(c.Param("articleId"))

// 	var article articles.Article
// 	if err := c.BindJSON(&article); err != nil {
// 		log.Println(err)
// 		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Unable to process entities"})
// 		return
// 	}

// 	res, err := UpdateArticle(article, authorId, articleId)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article updated successfully", "article": res})
// }

// func ArticleDelete(c *gin.Context) {
// 	authorId, _ := primitive.ObjectIDFromHex(c.Param("id"))
// 	articleId, _ := primitive.ObjectIDFromHex(c.Param("articleId"))

// 	res, err := DeleteArticle(authorId, articleId)

// 	if err != nil {
// 		log.Println(err)
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article deleted successfully", "article": res})
// }
