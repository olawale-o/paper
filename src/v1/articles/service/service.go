package service

import (
	"go-simple-rest/src/v1/articles/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	GetAll(params model.QueryParams) ([]model.Article, error)
	GetArticle(articleId primitive.ObjectID) (interface{}, error)
	Update(articleId primitive.ObjectID, article model.Article) (interface{}, error)
}
