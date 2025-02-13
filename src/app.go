package src

import (
	"go-simple-rest/src/v1/articles"
	"go-simple-rest/src/v1/auth/route"
	"go-simple-rest/src/v1/auth/transport"
	"go-simple-rest/src/v1/authors"
	"go-simple-rest/src/v1/comment"
	"go-simple-rest/src/v1/middlewares"

	"github.com/go-kit/log"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, svcEndpoint transport.Endpoints, logger log.Logger) {
	v1 := r.Group("/api/v1")
	{
		route.AuthRoutes(v1, svcEndpoint, logger)
		protected := v1.Group("/")
		protected.Use(middlewares.Auth())
		articles.ArticleRoutes(protected)
		authors.AuthorRoutes(protected)
		comment.CommentRoutes(protected)
	}
}
