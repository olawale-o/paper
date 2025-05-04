package repository

import (
	"go-simple-rest/src/v1/articles/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repostiory interface {
	GetArticles(filter, sort bson.M) ([]model.Article, error)
	GetArticleById(articleId primitive.ObjectID) (model.Article, error)
	UpdateArticle(filter, update bson.M) (any, error)
}
