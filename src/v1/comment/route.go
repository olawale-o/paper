package comment

import (
	"go-simple-rest/src/v1/comment/controller"
	"go-simple-rest/src/v1/comment/model"
	"go-simple-rest/src/v1/middlewares"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CommentRoutes(r *gin.RouterGroup, dbClient *mongo.Database) {
	controller := controller.CommentControllerImpl(dbClient)

	comments := r.Group("/articles/:id/comments")
	comments.GET("/", controller.Index)
	comments.GET("/:cid", controller.Show)
	comments.POST("/", middlewares.RequestToJSON[model.Comment](), middlewares.Validator[model.Comment](), controller.New)
	comments.POST("/:cid/reply", middlewares.RequestToJSON[model.Comment](), middlewares.Validator[model.Comment](), controller.ReplyComment)
}
