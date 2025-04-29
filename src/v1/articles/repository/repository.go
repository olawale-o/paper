package repository

import (
	"context"
	"go-simple-rest/src/v1/articles/model"

	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	Find(ctx context.Context, collection string, filter bson.M, sort bson.M, opts bson.M) ([]model.Article, error)
	FindOne(ctx context.Context, collection string, filter bson.M, opts bson.M) (model.Article, error)
	InsertOne(ctx context.Context, collection string, doc interface{}) (interface{}, error)
	FindOneAndUpdate(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (interface{}, error)
	DeleteOne(ctx context.Context, collection string, filter bson.M) error
	UpdateOne(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (interface{}, error)
}
