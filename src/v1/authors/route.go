package authors

import "github.com/gin-gonic/gin"

func AuthorRoutes(r *gin.RouterGroup) {
	r.POST("/author/articles", CreateArticle)
	r.PUT("/author/articles/:articleId", UpdateArticle)
	r.DELETE("/author/articles/:articleId", DeleteArticle)
}
