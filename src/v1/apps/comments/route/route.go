package route

import (
	"comments/controller"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.Engine) {

	// comments := v1.Group("articles/:id/comments")
	r.GET("/articles/:id/comments", controller.Index)
	r.GET("/articles/:id/comments/:cid", controller.Show)
	r.POST("/articles/:id/comments", controller.New)
}
