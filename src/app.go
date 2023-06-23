package src

import (
	"go-simple-rest/src/v1/articles"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	articles.ArticleRoutes(r)
}
