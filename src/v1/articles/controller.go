package articles

import (
	"go-simple-rest/src/v1/articles/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
func GetArticles(c *gin.Context) {
	date := c.DefaultQuery("date", "desc")
	likes := c.DefaultQuery("likes", "desc")
	views := c.DefaultQuery("views", "desc")
	params := model.QueryParams{
		Date:  date,
		Likes: likes,
		Views: views,
	}
	articles, err := GetAll(params)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, articles)
}

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
func ShowArticle(c *gin.Context) {
	article, err := GetArticle(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"article": article})
}

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
func UpdateArticle(c *gin.Context) {
	result, err := Update(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article updated", "article": result})
}
