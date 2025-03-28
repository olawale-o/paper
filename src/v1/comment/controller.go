package comment

import (
	"fmt"
	"go-simple-rest/db"
	"go-simple-rest/src/v1/comment/model"
	"go-simple-rest/src/v1/comment/repo"
	"go-simple-rest/src/v1/comment/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client, ctx, err = db.Connect()
var database = client.Database("go")

// Comment godoc
// @Summary Create a new comment
// @Description Create a new comment for an article
// @Tags articles
// @Accept json
// @Produce json
// @Param id path string true "Article ID"
// @Param comment body model.Comment true "Comment details"
// @Success 200 {object} string "Comment saved"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /articles/{id}/comments [post]
func New(c *gin.Context) {

	repository, err := repo.New(database)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	commentService, err := service.New(repository)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	articleId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	userId, _ := primitive.ObjectIDFromHex(c.MustGet("userId").(string))
	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Please provide valid credntials"})
		return
	}
	err, _ = commentService.NewComment(comment, articleId, userId)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Comment saved"})
}

func Show(c *gin.Context) {
	articleId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	commentId, _ := primitive.ObjectIDFromHex(c.Param("cid"))

	repository, err := repo.New(database)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	commentService, err := service.New(repository)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	var next primitive.ObjectID
	nextCursor := c.Query("nextCursor")

	if nextCursor == "" {
		next = primitive.NilObjectID
	} else {
		next, err = primitive.ObjectIDFromHex(nextCursor)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}

	res, err := commentService.GetComment(articleId, commentId, next)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "comment", "data": res})
}

// Comment godoc
// @Tags Articles
// @Summary Get article comments
// @Description Retrieves comments for a specific article.
// @Param id path string true "Article ID"
// @Param limit query int true "Limit"
// @Param prev query int true "Prev"
// @Param next query int true "Next"
// @Produce json
// @Success 200 {object} string "Response"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /articles/{id}/comments [get]
func Index(c *gin.Context) {
	repository, err := repo.New(database)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	commentService, err := service.New(repository)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// var limit int
	articleId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	// limit, err = strconv.Atoi(c.Query("limit"))
	// if err != nil {
	// 	limit = 0
	// }

	// prev := c.Query("prev")
	// next := c.Query("next")

	// res, err := commentService.GetComments(articleId, limit, prev, next)

	var next primitive.ObjectID
	nextCursor := c.Query("nextCursor")

	if nextCursor == "" {
		next = primitive.NilObjectID
	} else {
		next, err = primitive.ObjectIDFromHex(nextCursor)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
	}

	res, nextId, err := commentService.ArticleComments(articleId, next)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "comments", "nextCursor": nextId, "data": res})
}

// ReplyComment godoc
// @Tags Articles
// @Summary Reply to a comment
// @Description Replies to a specific comment.
// @Param id path string true "Article ID"
// @Param cid path string true "Comment ID"
// @Param content body string true "Content"
// @Produce json
// @Success 200 {object} string "Response"
// @Failure 400 {object} string "Error"
// @Failure 500 {object} string "Error"
// @Router /articles/{id}/comments/{cid}/reply [post]
func ReplyComment(c *gin.Context) {
	repository, err := repo.New(database)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	commentService, err := service.New(repository)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	articleId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	commentId, _ := primitive.ObjectIDFromHex(c.Param("cid"))
	userId, _ := primitive.ObjectIDFromHex(c.MustGet("userId").(string))

	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Please provide valid credntials"})
		return
	}

	_, err = commentService.ReplyComment(comment, articleId, commentId, userId)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Comment Saved"})
}
