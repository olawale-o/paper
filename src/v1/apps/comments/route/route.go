package route

import (
	"comments/controller"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.Engine) {
	r.GET("/", controller.Index)
	r.GET("/:cid", controller.Show)
	r.POST("/app-event", controller.Event)
}
