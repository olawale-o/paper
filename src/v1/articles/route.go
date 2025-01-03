package articles

import (
	"github.com/gin-gonic/gin"
)

func ArticleRoutes(r *gin.RouterGroup) {
	r.GET("/articles", GetArticles)
	r.POST("/articles", NewArticle)
	r.GET("/articles/:id", ShowArticle)
	r.PUT("/articles/:id", UpdateArticle)
	r.DELETE("/articles/:id", DeleteArticle)
}
