package repo

import (
	"context"
	"go-simple-rest/src/v1/comment/model"

	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	Get(ctx context.Context, collection string, filter bson.M, sort bson.M, limit int64) ([]model.Comment, error)
	FindOne(ctx context.Context, collection string, filter bson.M, v bson.M, opts bson.M) (interface{}, error)
	InsertOne(ctx context.Context, collection string, doc interface{}) (interface{}, error)
	// FindOneAndUpdate(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (interface{}, error)
	// DeleteOne(ctx context.Context, collection string, filter bson.M) error
	UpdateOne(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (interface{}, error)
	Aggregate(ctx context.Context, collection string, pipeline []bson.M) ([]model.ArticleWithComments, error)
}
