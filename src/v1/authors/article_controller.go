package authors

import (
	"go-simple-rest/db"
	"go-simple-rest/src/v1/authors/model"
	"go-simple-rest/src/v1/authors/repo"
	"go-simple-rest/src/v1/authors/service"
	"go-simple-rest/src/v1/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client, ctx, err = db.Connect()

var database = client.Database("go")

// AuthorArticleIndex godoc
// @Tags Authors
// @Summary Get all articles by author
// @Description Retrieves all articles written by a specific author.
// @Param id path string true "Author ID"
// @Produce json
// @Success 200 {object} string "Response"
// @Failure 400 {object} string "Error"
// @Failure 500 {object} string "Error"
// @Router /authors/{id}/articles [get]
func ArticleIndex(c *gin.Context) {
	oid, err := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	repository, err := repo.New(database)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	service, err := service.New(repository)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	res, _ := service.AllArticles(oid)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article retrieved successfully", "articles": res})
}

// AuthorArticleNew godoc
// @Tags Authors
// @Summary Create a new article written by a specific author
// @Description Creates a new article written by a specific author.
// @Param id path string true "Author ID"
// @Param article body model.AuthorArticle true "Article"
// @Produce json
// Accept application/json
// @Success 201 {object} string "Response"
// @Failure 400 {object} string "Error"
// @Failure 500 {object} string "Error"
// @Router /authors/{id}/articles [post]
func ArticleNew(c *gin.Context) {
	oid, err := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	repository, err := repo.New(database)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	service, err := service.New(repository)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	var newArticle model.AuthorArticle
	if err := c.BindJSON(&newArticle); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Unable to process entities"})
		return
	}
	res, err := service.CreateArticle(newArticle, oid)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article created successfully", "article": res})
}

// AuthorArticleNew godoc
// @Tags Authors
// @Summary Update an article written by a specific author
// @Description Updates an article written by a specific author.
// @Param id path string true "Author ID"
// @Param articleId path string true "Article ID"
// @Param article body model.AuthorArticleUpdateRequest true "Article"
// @Produce json
// Accept application/json
// @Success 200 {object} model.AuthorArticleUpdateResponse "Response"
// @Failure 400 {object} string "Error"
// @Failure 500 {object} string "Error"
// @Router /authors/{id}/articles/{articleId} [put]
func ArticleUpdate(c *gin.Context) {
	authorId, err := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid author ID"})
		return
	}
	articleId, err := utils.ParseParamToPrimitiveObjectId(c.Param("articleId"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid article ID"})
		return
	}

	var article model.AuthorArticle
	if err := c.BindJSON(&article); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Unable to process entities"})
		return
	}

	repository, err := repo.New(database)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	service, err := service.New(repository)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	var newArticle model.AuthorArticle
	if err := c.BindJSON(&newArticle); err != nil {
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

// AuthorArticleNew godoc
// @Tags Authors
// @Summary Delete an article written by a specific author
// @Description Deletes an article written by a specific author.
// @Param id path string true "Author ID"
// @Param articleId path string true "Article ID"
// @Produce json
// Accept application/json
// @Success 200 {object} model.AuthorArticleUpdateResponse "Response"
// @Failure 400 {object} string "Error"
// @Failure 500 {object} string "Error"
// @Router /authors/{id}/articles/{articleId} [delete]
func ArticleDelete(c *gin.Context) {
	authorId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	articleId, _ := primitive.ObjectIDFromHex(c.Param("articleId"))

	repository, err := repo.New(database)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	service, err := service.New(repository)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	var newArticle model.AuthorArticle
	if err := c.BindJSON(&newArticle); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Unable to process entities"})
		return
	}

	err = service.DeleteArticle(authorId, articleId)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
