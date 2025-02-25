package repo

import (
	"context"
	"fmt"
	"go-simple-rest/db"
	"go-simple-rest/src/v1/comment/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client, ctx, err = db.Connect()

var collection = client.Database("go").Collection("articles")

type repository struct {
	db     *mongo.Database
	logger log.Logger
}

func New(db *mongo.Database) (model.Repository, error) {
	return &repository{
		db: db,
	}, nil
}

func (repo *repository) Get(ctx context.Context, collection string, filter bson.M, sort bson.M, limit int64) ([]model.Comment, error) {
	opts := options.Find().SetSort(sort).SetLimit(limit)
	cursor, err := repo.db.Collection(collection).Find(context.TODO(), filter, opts)

	if err != nil {
		fmt.Println(err.Error())
		return []model.Comment{}, err
	}
	var comments []model.Comment
	if err = cursor.All(context.TODO(), &comments); err != nil {
		return []model.Comment{}, err
	}
	return comments, nil
}

func (repo *repository) FindOne(ctx context.Context, collection string, filter bson.M, v bson.M) (interface{}, error) {
	if err := repo.db.Collection(collection).FindOne(context.TODO(), filter).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}
