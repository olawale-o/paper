package articles

import (
	"go-simple-rest/src/v1/articles/controller"
	"go-simple-rest/src/v1/articles/model"
	"go-simple-rest/src/v1/middlewares"

	"github.com/gin-gonic/gin"
)

func ArticleRoutes(r *gin.RouterGroup) {

	articles := r.Group("/articles")
	articles.GET("/", controller.GetArticles)
	articles.GET("/:id", controller.ShowArticle)
	articles.PUT("/:id", middlewares.RequestToJSON[model.Article](), middlewares.Validator[model.Article](), controller.UpdateArticle)
}
