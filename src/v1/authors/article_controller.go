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

var client, _, _ = db.Connect()

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
	repository, _ := repo.New(database)
	service, _ := service.New(repository)

	oid, err := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusBadRequest, Success: false, Message: "Invalid ID", Data: nil})
		return
	}
	res, _ := service.AllArticles(oid)
	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusOK, Success: true, Message: "Retrieved all articles", Data: res})
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
	repository, _ := repo.New(database)
	service, _ := service.New(repository)

	oid, err := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusBadRequest, Success: false, Message: "Invalid ID", Data: nil})
		return
	}

	newArticle := c.MustGet("body").(model.AuthorArticle)
	res, err := service.CreateArticle(newArticle, oid)

	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusInternalServerError, Success: false, Message: err.Error(), Data: nil})

		return
	}
	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusCreated, Success: true, Message: "Article created successfully", Data: res})
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
	repository, _ := repo.New(database)
	service, _ := service.New(repository)
	article := c.MustGet("body").(model.AuthorArticle)

	authorId, err := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusBadRequest, Success: false, Message: "Invalid ID", Data: nil})
		return
	}
	articleId, err := utils.ParseParamToPrimitiveObjectId(c.Param("articleId"))
	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusBadRequest, Success: false, Message: "Invalid ID", Data: nil})
		return
	}

	res, err := service.UpdateArticle(article, authorId, articleId)

	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusInternalServerError, Success: false, Message: err.Error(), Data: nil})
		return
	}

	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusOK, Success: true, Message: "Article updated successfully", Data: res})
}

// AuthorArticleDelete godoc
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

	repository, _ := repo.New(database)
	service, _ := service.New(repository)

	err := service.DeleteArticle(authorId, articleId)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
