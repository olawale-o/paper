package authors

import (
	"go-simple-rest/src/v1/authors/controller"
	"go-simple-rest/src/v1/authors/model"
	"go-simple-rest/src/v1/middlewares"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthorRoutes(r *gin.RouterGroup, databaseClient *mongo.Database) {

	c := controller.AuthorControllerImpl(databaseClient)
	articleController := controller.AuthorArticleControllerImpl(databaseClient)
	authors := r.Group("/authors")

	authors.GET("/:id", c.Show)
	authors.PUT("/:id", middlewares.RequestToJSON[model.Author](), middlewares.Validator[model.Author](), c.Update)
	authors.DELETE("/:id", c.Delete)
	authors.GET("/:id/articles", articleController.ArticleIndex)
	authors.POST("/:id/articles", middlewares.RequestToJSON[model.AuthorArticle](), middlewares.Validator[model.AuthorArticle](), articleController.ArticleNew)
	authors.PUT("/:id/articles/:articleId", middlewares.RequestToJSON[model.AuthorArticle](), middlewares.Validator[model.AuthorArticle](), articleController.ArticleUpdate)
	authors.DELETE("/:id/articles/:articleId", articleController.ArticleDelete)
}
