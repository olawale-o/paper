package controller

import (
	"go-simple-rest/src/v1/comment/dao"
	"go-simple-rest/src/v1/comment/model"
	"go-simple-rest/src/v1/comment/repo"
	"go-simple-rest/src/v1/comment/service"
	"go-simple-rest/src/v1/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentController struct {
	database *mongo.Database
}

func CommentControllerImpl(database *mongo.Database) Controller {
	return &CommentController{
		database: database,
	}
}

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
func (comnentController *CommentController) New(c *gin.Context) {

	commentDAO, _ := dao.New(comnentController.database)
	repository := repo.NewRepository(commentDAO)
	commentService, _ := service.New(repository)

	articleId, _ := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	userId, _ := utils.ParseParamToPrimitiveObjectId(c.MustGet("userId").(string))

	comment := c.MustGet("body").(model.Comment)
	err, _ := commentService.NewComment(comment, articleId, userId)

	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusInternalServerError, Success: false, Message: "Failed to create comment", Data: nil})
		return
	}
	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusOK, Success: true, Message: "Comment saved", Data: nil})
}

func (comnentController *CommentController) Show(c *gin.Context) {
	articleId, _ := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	commentId, _ := utils.ParseParamToPrimitiveObjectId(c.Param("cid"))

	commentDAO, _ := dao.New(comnentController.database)
	repository := repo.NewRepository(commentDAO)
	commentService, _ := service.New(repository)

	var next primitive.ObjectID
	var err error
	nextCursor := c.Query("nextCursor")

	if nextCursor == "" {
		next = primitive.NilObjectID
	} else {
		next, err = utils.ParseParamToPrimitiveObjectId(nextCursor)
		if err != nil {
			utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusBadRequest, Success: false, Message: "Invalid ID", Data: nil})
			return
		}
	}

	res, err := commentService.GetComment(articleId, commentId, next)

	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusNotFound, Success: false, Message: "Comment not found", Data: nil})
		return
	}
	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusOK, Success: true, Message: "Comment retrieved", Data: res})
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
func (comnentController *CommentController) Index(c *gin.Context) {
	commentDAO, _ := dao.New(comnentController.database)
	repository := repo.NewRepository(commentDAO)
	commentService, _ := service.New(repository)

	articleId, _ := utils.ParseParamToPrimitiveObjectId(c.Param("id"))

	var next primitive.ObjectID
	var err error
	nextCursor := c.Query("nextCursor")

	if nextCursor == "" {
		next = primitive.NilObjectID
	} else {
		next, err = utils.ParseParamToPrimitiveObjectId(nextCursor)
		if err != nil {
			utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusBadRequest, Success: false, Message: "Invalid ID", Data: nil})
			return
		}
	}

	res, nextId, err := commentService.ArticleComments(articleId, next)
	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusInternalServerError, Success: false, Message: "Failed to retrieve comments", Data: nil})
		return
	}
	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusOK, Success: true, Message: "Comments retrieved", Data: map[string]interface{}{"comments": res, "nextCursor": nextId}})
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
func (comnentController *CommentController) ReplyComment(c *gin.Context) {
	commentDAO, _ := dao.New(comnentController.database)
	repository := repo.NewRepository(commentDAO)
	commentService, _ := service.New(repository)

	articleId, _ := utils.ParseParamToPrimitiveObjectId(c.Param("id"))
	commentId, _ := utils.ParseParamToPrimitiveObjectId(c.Param("cid"))
	userId, _ := utils.ParseParamToPrimitiveObjectId(c.MustGet("userId").(string))

	comment := c.MustGet("body").(model.Comment)

	_, err := commentService.ReplyComment(comment, articleId, commentId, userId)
	if err != nil {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusInternalServerError, Success: false, Message: err.Error(), Data: nil})
		return
	}
	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusOK, Success: true, Message: "Comment saved", Data: nil})
}
