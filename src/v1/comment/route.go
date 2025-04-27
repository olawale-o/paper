package comment

import (
	"go-simple-rest/src/v1/comment/model"
	"go-simple-rest/src/v1/middlewares"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(r *gin.RouterGroup) {

	comments := r.Group("/articles/:id/comments")
	comments.GET("/", Index)
	comments.GET("/:cid", Show)
	comments.POST("/", middlewares.RequestToJSON[model.Comment](), middlewares.Validator[model.Comment](), New)
	comments.POST("/:cid/reply", ReplyComment)
}
