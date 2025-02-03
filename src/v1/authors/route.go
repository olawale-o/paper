package authors

import "github.com/gin-gonic/gin"

func AuthorRoutes(r *gin.RouterGroup) {
	authors := r.Group("/authors")
	authors.GET("/", Index)
	authors.GET("/:id", Show)
	authors.PUT("/:id", Update)
	authors.DELETE("/:id", Delete)
	authors.GET("/:id/articles", ArticleIndex)
	authors.POST("/:id/articles", ArticleNew)
	authors.PUT("/:id/articles/:articleId", ArticleUpdate)
	authors.DELETE("/:id/articles/:articleId", ArticleDelete)
}
