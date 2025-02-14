package route

import (
	"articles/controller"

	"github.com/gin-gonic/gin"
)

func ArticleRoutes(r *gin.Engine) {
	// articles := r.Group("/articles")
	// articles.Use(middlewares.Auth())
	r.GET("/", controller.GetArticles)
	r.POST("/", controller.NewArticle)
	r.GET("/:id", controller.ShowArticle)
	r.PUT("/:id", controller.UpdateArticle)
	r.DELETE("/:id", controller.DeleteArticle)
}
