package controller

import (
	"github.com/gin-gonic/gin"
)

type AuthorArticleController interface {
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
	ArticleIndex(c *gin.Context)

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
	ArticleNew(c *gin.Context)

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
	ArticleUpdate(c *gin.Context)

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
	ArticleDelete(c *gin.Context)
}
