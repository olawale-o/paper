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

var collection = client.Database("go").Collection("comments")

type RepositoryManager struct {
	db     *mongo.Database
	logger log.Logger
}

func New(db *mongo.Database) (Repository, error) {
	return &RepositoryManager{
		db: db,
	}, nil
}

func (repo *RepositoryManager) Get(ctx context.Context, collection string, filter bson.M, sort bson.M, limit int64) ([]model.Comment, error) {
	// options := options.Find().SetSort(sort).SetProjection(opts)
	cursor, err := repo.db.Collection(collection).Find(context.TODO(), filter)

	if err != nil {
		fmt.Println(err.Error())
		return []model.Comment{}, err
	}
	var articles []model.Comment
	if err = cursor.All(context.TODO(), &articles); err != nil {
		return []model.Comment{}, err
	}
	return articles, nil
}

func (repo *RepositoryManager) FindOne(ctx context.Context, collection string, filter bson.M, v bson.M, opts bson.M) (interface{}, error) {
	options := options.FindOne().SetProjection(opts)
	if err := repo.db.Collection(collection).FindOne(context.TODO(), filter, options).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

func (repo *RepositoryManager) InsertOne(ctx context.Context, collection string, doc interface{}) (interface{}, error) {
	result, err := repo.db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}
