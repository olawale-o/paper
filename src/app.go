package src

import (
	"go-simple-rest/src/v1/articles"
	"go-simple-rest/src/v1/auth"
	"go-simple-rest/src/v1/authors"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	articles.ArticleRoutes(r)
	auth.AuthRoutes(r)
	authors.AuthorRoutes(r)
}
