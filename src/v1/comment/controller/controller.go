package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
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
	New(c *gin.Context)

	// Comment godoc
	// @Summary Show a comment
	// @Description Show a comment for an article
	// @Tags articles
	// @Accept json
	// @Produce json
	// @Param id path string true "Article ID"
	// @Param cid path string true "Comment ID"
	// @Success 200 {object} string "Comment saved"
	// @Failure 400 {object} string "Bad Request"
	// @Failure 500 {object} string "Internal Server Error"
	// @Router /articles/{id}/comments/{cid} [get]
	Show(c *gin.Context)

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
	Index(c *gin.Context)
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
	ReplyComment(c *gin.Context)
}
