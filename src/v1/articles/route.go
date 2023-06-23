package articles

import (
	"github.com/gin-gonic/gin"
)

func ArticleRoutes(r *gin.Engine) {
	r.GET("/articles", GetArticles)
}
