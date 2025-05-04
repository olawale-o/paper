// package repo

// import (
// 	"context"
// 	"go-simple-rest/src/v1/authors/model"

// 	"go.mongodb.org/mongo-driver/bson"
// )

// type Repository interface {
// 	Get(ctx context.Context, collection string, filter bson.M) ([]model.AuthorArticle, error)
// 	FindOne(ctx context.Context, collection string, filter bson.M, v model.Author) (model.Author, error)
// 	InsertOne(ctx context.Context, collection string, doc any) (any, error)
// 	FindOneAndUpdate(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (model.AuthorArticleUpdateResponse, error)
// 	DeleteOne(ctx context.Context, collection string, filter bson.M) error
// 	UpdateOne(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (any, error)
// }

package repo

import (
	"go-simple-rest/src/v1/authors/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	UpdateArticleAuthor(filter, update bson.M) (model.AuthorArticleUpdateResponse, error)
	GetAuthorArticles(authorId primitive.ObjectID) ([]model.AuthorArticle, error)
	Create(doc model.AuthorArticle) (any, error)
	Update(filter, update bson.M) (model.AuthorArticleUpdateResponse, error)
	Delete(articleId primitive.ObjectID) error
	GetAuthorById(filter bson.M) (any, error)
	UpdateAuthor(filter, update bson.M) (any, error)
	DeleteAuthor(filter bson.M) error
}
