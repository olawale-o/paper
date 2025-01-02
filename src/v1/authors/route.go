package authors

import "github.com/gin-gonic/gin"

func AuthorRoutes(r *gin.Engine) {
	r.POST("/author/:authorId/articles", CreateArticle)
	r.PUT("/author/:authorId/articles/:articleId", UpdateArticle)
	r.DELETE("/author/:authorId/articles/:articleId", DeleteArticle)
}
