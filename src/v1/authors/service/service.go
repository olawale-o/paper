package service

import (
	"go-simple-rest/src/v1/authors/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	AllArticles(authorId primitive.ObjectID) (any, error)
	CreateArticle(article model.AuthorArticle, authorId primitive.ObjectID) (any, error)
	UpdateArticle(article model.AuthorArticle, authorId primitive.ObjectID, articleId primitive.ObjectID) (any, error)
	DeleteArticle(authorId primitive.ObjectID, articleId primitive.ObjectID) error
	ShowAuthor(authorId primitive.ObjectID) (model.Author, error)
	UpdateAuthor(authorId primitive.ObjectID, updatedAuthor model.Author) (any, error)
	DeleteAuthor(authorId primitive.ObjectID) (int64, error)
}
