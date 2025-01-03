package src

import (
	"go-simple-rest/src/v1/articles"
	"go-simple-rest/src/v1/auth"
	"go-simple-rest/src/v1/authors"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		articles.ArticleRoutes(v1)
		auth.AuthRoutes(v1)
		authors.AuthorRoutes(v1)
	}
}
