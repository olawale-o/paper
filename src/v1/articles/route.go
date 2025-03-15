package articles

import (
	"go-simple-rest/src/v1/articles/controller"

	"github.com/gin-gonic/gin"
)

func ArticleRoutes(r *gin.RouterGroup) {

	articles := r.Group("/articles")
	articles.GET("/", controller.GetArticles)
	articles.GET("/:id", controller.ShowArticle)
	articles.PUT("/:id", controller.UpdateArticle)
}
