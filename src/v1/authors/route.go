package authors

import (
	"go-simple-rest/src/v1/authors/model"
	"go-simple-rest/src/v1/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthorRoutes(r *gin.RouterGroup) {
	authors := r.Group("/authors")
	// authors.GET("/", Index)
	authors.GET("/:id", Show)
	authors.PUT("/:id", middlewares.RequestToJSON[model.Author](), middlewares.Validator[model.Author](), Update)
	authors.DELETE("/:id", Delete)
	authors.GET("/:id/articles", ArticleIndex)
	authors.POST("/:id/articles", middlewares.RequestToJSON[model.AuthorArticle](), middlewares.Validator[model.AuthorArticle](), ArticleNew)
	authors.PUT("/:id/articles/:articleId", middlewares.RequestToJSON[model.AuthorArticle](), middlewares.Validator[model.AuthorArticle](), ArticleUpdate)
	authors.DELETE("/:id/articles/:articleId", ArticleDelete)
}
