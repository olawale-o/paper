package route

import (
	"articles/controller"
	"articles/middlewares"

	"github.com/gin-gonic/gin"
)

func ArticleRoutes(r *gin.Engine) {
	// articles := r.Group("/articles")
	r.Use(middlewares.Auth())
	r.GET("/", controller.GetArticles)
	r.POST("/", controller.NewArticle)
	r.GET("/:id", controller.ShowArticle)
	r.PUT("/:id", controller.UpdateArticle)
	r.DELETE("/:id", controller.DeleteArticle)
	// r.GET("/:id/comments", controller.Index)
	// r.GET("/:id/comments/:cid", controller.Show)
	r.POST("/:id/comments", controller.NewComment)
}
