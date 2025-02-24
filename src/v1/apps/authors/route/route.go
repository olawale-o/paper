package route

import (
	"authors/controller"

	"github.com/gin-gonic/gin"
)

func AuthorRoutes(r *gin.Engine) {

	r.GET("/", controller.Index)
	r.GET("/:id", controller.Show)
	r.PUT("/:id", controller.Update)
	r.DELETE("/:id", controller.Delete)
	r.GET("/:id/articles", controller.ArticleIndex)
	// r.POST("/:id/articles", controller.ArticleNew)
	r.PUT("/:id/articles/:articleId", controller.ArticleUpdate)
	r.DELETE("/:id/articles/:articleId", controller.ArticleDelete)
	r.POST("/app-event", controller.Event)

}
