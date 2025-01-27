package authors

import "github.com/gin-gonic/gin"

func AuthorRoutes(r *gin.RouterGroup) {
	r.GET("/authors", Index)
	r.GET("/authors/:id", Show)
	r.PUT("/authors/:id", Update)
	r.DELETE("/authors/:id", Delete)
	r.GET("/authors/:id/articles", ArticleIndex)
	r.POST("/authors/:id/articles", ArticleNew)
	r.PUT("/authors/:id/articles/:articleId", ArticleUpdate)
	r.DELETE("/authors/:id/articles/:articleId", ArticleDelete)
}
