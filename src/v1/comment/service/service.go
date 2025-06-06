package service

import (
	"go-simple-rest/src/v1/comment/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Comments   []model.Comment `json:"comments"`
	HasNext    bool            `json:"hasNext"`
	HasPrev    bool            `json:"hasPrev"`
	PreviousID string          `json:"previousId,omitempty"`
	NextID     string          `json:"nextId,omitempty"`
}

type Service interface {
	NewComment(c model.Comment, articleId primitive.ObjectID, userId primitive.ObjectID) (error, any)
	GetComment(articleId primitive.ObjectID, commentId primitive.ObjectID, next primitive.ObjectID) (any, error)
	GetComments(articleId primitive.ObjectID, l int, prev string, next string) (Response, error)
	ReplyComment(c model.Comment, articleId primitive.ObjectID, commentId primitive.ObjectID, userId primitive.ObjectID) (any, error)
	ArticleComments(articleId primitive.ObjectID, next primitive.ObjectID) ([]model.ArticleWithComments, any, error)
	MoreReplies(articleId primitive.ObjectID, commentId primitive.ObjectID, next primitive.ObjectID) ([]model.Comment, error)
}
