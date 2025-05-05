package controller

import "github.com/gin-gonic/gin"

type Controller interface {

	// Articles godoc
	// @Tags Articles
	// @Summary Get articles
	// @Description Retrieves articles
	// @Param date query string false "Sort by date" example("desc")
	// @Param likes query string false "Sort by likes" example("desc")
	// @Param views query string false "Sort by views"
	// @Produce json
	// @Success 200 {object} []model.Article "Response"
	// @Failure 400 {object} string "Error"
	// @Failure 500 {object} string "Error"
	// @Router /articles [get]
	GetArticles(c *gin.Context)

	// Articles godoc
	// @Tags Articles
	// @Summary Get articles by id
	// @Description Retrieves a specific article by ID.
	// @Param id path string true "Article ID"
	// @Produce json
	// @Success 200 {object} model.Article "Response"
	// @Failure 400 {object} string "Error"
	// @Failure 500 {object} string "Error"
	// @Router /articles/{id} [get]
	ShowArticle(c *gin.Context)

	// Articles godoc
	// @Tags Articles
	// @Summary Update articles by id
	// @Description Updates a specific article by ID.
	// @Param id path string true "Article ID"
	// @Produce json
	// @Success 200 {object} string "Response"
	// @Failure 400 {object} string "Error"
	// @Failure 500 {object} string "Error"
	// @Router /articles/{id} [put]
	UpdateArticle(c *gin.Context)
}
