package dao

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

type MongoDBCommentDaoManager struct {
	db     *mongo.Database
	logger log.Logger
}

func New(db *mongo.Database) (CommentDAO, error) {
	return &MongoDBCommentDaoManager{
		db: db,
	}, nil
}

func (repo *MongoDBCommentDaoManager) Find(ctx context.Context, collection string, filter bson.M, sort bson.M, limit int64) ([]model.Comment, error) {
	options := options.Find().SetSort(sort)
	// .SetProjection(opts)
	cursor, err := repo.db.Collection(collection).Find(context.TODO(), filter, options)

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

func (repo *MongoDBCommentDaoManager) FindOne(ctx context.Context, collection string, filter bson.M, v bson.M, opts bson.M) (any, error) {
	options := options.FindOne().SetProjection(opts)
	if err := repo.db.Collection(collection).FindOne(context.TODO(), filter, options).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

func (repo *MongoDBCommentDaoManager) InsertOne(ctx context.Context, collection string, doc any) (any, error) {
	result, err := repo.db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (repo *MongoDBCommentDaoManager) UpdateOne(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (any, error) {
	result, err := repo.db.Collection(collection).UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(upsert))
	if err != nil {
		return nil, err
	}
	return result.UpsertedID, nil
}

func (repo *MongoDBCommentDaoManager) Aggregate(ctx context.Context, collection string, pipeline []bson.M) ([]model.ArticleWithComments, error) {
	cursor, err := repo.db.Collection(collection).Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	var results []model.ArticleWithComments
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
