package articles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	TITLE  string `json:"title"`
	AUTHOR string `json:"author"`
}

var a []Article = []Article{
	{TITLE: "Article 1", AUTHOR: "Author 1"},
	{TITLE: "Article 2", AUTHOR: "Author 2"},
	{TITLE: "Article 3", AUTHOR: "Author 3"},
	{TITLE: "Article 4", AUTHOR: "Author 4"},
}

func GetArticles(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	//  "articles": articles,
	// })
	c.IndentedJSON(http.StatusOK, a)
}
