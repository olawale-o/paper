package articles

import (
	"github.com/gin-gonic/gin"
)

func ArticleRoutes(r *gin.RouterGroup) {

	articles := r.Group("/articles")
	// articles.GET("/", GetArticles)
	articles.GET("/:id", ShowArticle)
	articles.PUT("/:id", UpdateArticle)
}
