package comment

import (
	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.RouterGroup) {

	comments := r.Group("/articles/:id/comments")
	comments.GET("/", Index)
	comments.GET("/:cid", Show)
	comments.POST("/", New)
}
