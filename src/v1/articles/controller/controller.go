package controller

import (
	"go-simple-rest/db"
	"go-simple-rest/src/v1/articles/model"
	"go-simple-rest/src/v1/articles/repository/implementation"
	"go-simple-rest/src/v1/articles/service"
	"go-simple-rest/src/v1/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var client, ctx, err = db.Connect()
var collection = client.Database("go").Collection("articles")
var database = client.Database("go")

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
	repo, err := implementation.New(database)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	service, err := service.New(repo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	articles, err := service.GetAll(model.QueryParams{
		Date:  date,
		Likes: likes,
		Views: views,
	})
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
	repo, err := implementation.New(database)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	oid, err := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	service, err := service.New(repo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	article, err := service.GetArticle(oid)
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
	repo, err := implementation.New(database)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	oid, err := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var updatedArticle model.Article
	if err := c.BindJSON(&updatedArticle); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	service, err := service.New(repo)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	result, err := service.Update(oid, updatedArticle)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article updated", "article": result})
}
