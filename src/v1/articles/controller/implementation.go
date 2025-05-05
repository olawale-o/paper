package controller

import (
	"go-simple-rest/src/v1/articles/dao"
	"go-simple-rest/src/v1/articles/model"
	"go-simple-rest/src/v1/articles/repository"
	"go-simple-rest/src/v1/articles/service"
	"go-simple-rest/src/v1/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleController struct {
	database *mongo.Database
}

func ArticleControllerImpl(database *mongo.Database) Controller {
	return &ArticleController{
		database: database,
	}
}

func (controller *ArticleController) GetArticles(c *gin.Context) {
	date := c.DefaultQuery("date", "desc")
	likes := c.DefaultQuery("likes", "desc")
	views := c.DefaultQuery("views", "desc")
	articleDAO, _ := dao.New(controller.database)
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

func (controller *ArticleController) ShowArticle(c *gin.Context) {
	articleDAO, _ := dao.New(controller.database)
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

func (controller *ArticleController) UpdateArticle(c *gin.Context) {
	articleDAO, _ := dao.New(controller.database)
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
