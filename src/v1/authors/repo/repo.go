package repo

import (
	"context"
	"go-simple-rest/src/v1/authors/model"

	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	Get(ctx context.Context, collection string, filter bson.M) ([]model.AuthorArticle, error)
	FindOne(ctx context.Context, collection string, filter bson.M, v model.Author) (model.Author, error)
	InsertOne(ctx context.Context, collection string, doc interface{}) (interface{}, error)
	FindOneAndUpdate(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (model.AuthorArticleUpdateResponse, error)
	DeleteOne(ctx context.Context, collection string, filter bson.M) error
	UpdateOne(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (interface{}, error)
}
