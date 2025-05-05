package src

import (
	"go-simple-rest/src/v1/articles"
	"go-simple-rest/src/v1/auth/route"
	"go-simple-rest/src/v1/authors"
	"go-simple-rest/src/v1/comment"
	"go-simple-rest/src/v1/middlewares"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routes(r *gin.Engine, databaseClient *mongo.Database) {
	v1 := r.Group("/api/v1")
	{
		route.AuthRoutes(v1, databaseClient)
		protected := v1.Group("/")
		protected.Use(middlewares.Auth())
		articles.ArticleRoutes(protected, databaseClient)
		authors.AuthorRoutes(protected, databaseClient)
		comment.CommentRoutes(protected, databaseClient)
	}
}
