package repo

import (
	"go-simple-rest/src/v1/comment/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Create(comment model.Comment) (any, error)
	FindAll(filter, sort bson.M, limit int64) ([]model.Comment, error)
	FindById(filter bson.M) (model.Comment, error)
	FindArticleById(id primitive.ObjectID) (any, error)
	FindCommentByIdWithReplies(articleId, commentId, nextId primitive.ObjectID) ([]model.Comment, error)
	UpdateCommentWithReply(id, articleId primitive.ObjectID, update bson.M) (any, error)
	Aggregate(pipeline []bson.M) ([]model.ArticleWithComments, error)
}
