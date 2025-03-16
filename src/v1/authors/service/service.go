package service

import (
	"go-simple-rest/src/v1/authors/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	AllArticles(authorId primitive.ObjectID) (interface{}, error)
	CreateArticle(article model.AuthorArticle, authorId primitive.ObjectID) (interface{}, error)
	UpdateArticle(article model.AuthorArticle, authorId primitive.ObjectID, articleId primitive.ObjectID) (interface{}, error)
	DeleteArticle(authorId primitive.ObjectID, articleId primitive.ObjectID) error
	ShowAuthor(authorId primitive.ObjectID) (interface{}, error)
	UpdateAuthor(authorId primitive.ObjectID, updatedAuthor model.Author) (interface{}, error)
	DeleteAuthor(authorId primitive.ObjectID) (int64, error)
}
