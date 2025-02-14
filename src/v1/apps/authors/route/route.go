package route

import (
	"authors/controller"

	"github.com/gin-gonic/gin"
)

func AuthorRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1/articles")
	{
		v1.GET("/", controller.Index)
		v1.GET("/:id", controller.Show)
		v1.PUT("/:id", controller.Update)
		v1.DELETE("/:id", controller.Delete)
		authors := r.Group("/authors")
		authors.GET("/:id/articles", controller.ArticleIndex)
		authors.POST("/:id/articles", controller.ArticleNew)
		authors.PUT("/:id/articles/:articleId", controller.ArticleUpdate)
		authors.DELETE("/:id/articles/:articleId", controller.ArticleDelete)
	}
}
