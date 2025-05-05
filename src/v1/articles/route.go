package articles

import (
	"go-simple-rest/src/v1/articles/controller"
	"go-simple-rest/src/v1/articles/model"
	"go-simple-rest/src/v1/middlewares"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ArticleRoutes(r *gin.RouterGroup, databaseClient *mongo.Database) {

	controller := controller.ArticleControllerImpl(databaseClient)
	articles := r.Group("/articles")
	articles.GET("/", controller.GetArticles)
	articles.GET("/:id", controller.ShowArticle)
	articles.PUT("/:id", middlewares.RequestToJSON[model.Article](), middlewares.Validator[model.Article](), controller.UpdateArticle)
}
