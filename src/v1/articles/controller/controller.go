package controller

import (
	"go-simple-rest/db"
	"go-simple-rest/src/v1/articles/dao"
	"go-simple-rest/src/v1/articles/model"
	"go-simple-rest/src/v1/articles/repository"
	"go-simple-rest/src/v1/articles/service"
	"go-simple-rest/src/v1/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var client, _, _ = db.Connect()

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
	articleDAO, _ := dao.New(database)
	repo := repository.NewRepositoryManager(articleDAO)
	service, _ := service.New(repo)

	articles, err := service.GetAll(model.QueryParams{
		Date:  date,
		Likes: likes,
		Views: views,
	})
	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusInternalServerError, Message: err.Error(), Success: false, Data: nil})

		return
	}
	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusOK, Message: "articles", Success: true, Data: articles})
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
	articleDAO, _ := dao.New(database)
	repo := repository.NewRepositoryManager(articleDAO)
	service, _ := service.New(repo)
	oid, err := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusBadRequest, Message: err.Error(), Success: false, Data: nil})
		return
	}

	article, err := service.GetArticle(oid)
	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusNotFound, Message: err.Error(), Success: false, Data: nil})
		return
	}

	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusOK, Message: "articles", Success: true, Data: article})
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
	articleDAO, _ := dao.New(database)
	repo := repository.NewRepositoryManager(articleDAO)
	service, _ := service.New(repo)

	oid, err := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusBadRequest, Message: err.Error(), Success: false, Data: nil})
		return
	}

	updatedArticle := c.MustGet("body").(model.Article)

	result, err := service.Update(oid, updatedArticle)
	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusNotFound, Message: err.Error(), Success: false, Data: nil})
		return
	}
	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusOK, Message: "articles", Success: true, Data: result})

}
