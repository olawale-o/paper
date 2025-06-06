package dao

import (
	"context"
	"fmt"
	"go-simple-rest/db"
	"go-simple-rest/src/v1/articles/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client, ctx, err = db.Connect()

type MongoDBArticleDaoManager struct {
	db     *mongo.Database
	logger log.Logger
}

func New(db *mongo.Database) (ArticleDAO, error) {
	return &MongoDBArticleDaoManager{
		db: db,
	}, nil
}

func (repo *MongoDBArticleDaoManager) Find(ctx context.Context, collection string, filter bson.M, sort bson.M, opts bson.M) ([]model.Article, error) {
	options := options.Find().SetSort(sort).SetProjection(opts)
	cursor, err := repo.db.Collection(collection).Find(context.TODO(), filter, options)

	if err != nil {
		fmt.Println(err.Error())
		return []model.Article{}, err
	}
	var articles []model.Article
	if err = cursor.All(context.TODO(), &articles); err != nil {
		return []model.Article{}, err
	}
	return articles, nil
}

func (repo *MongoDBArticleDaoManager) FindOne(ctx context.Context, collection string, filter bson.M, opts bson.M) (model.Article, error) {
	var v model.Article
	options := options.FindOne().SetProjection(opts)
	if err := repo.db.Collection(collection).FindOne(context.TODO(), filter, options).Decode(&v); err != nil {
		return v, err
	}
	return v, nil
}

func (repo *MongoDBArticleDaoManager) InsertOne(ctx context.Context, collection string, doc any) (any, error) {
	res, err := repo.db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		return "", err
	}
	return res.InsertedID, nil
}

func (repo *MongoDBArticleDaoManager) FindOneAndUpdate(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (any, error) {
	var data any
	opts := options.FindOneAndUpdate().SetUpsert(upsert)
	repo.db.Collection(collection).FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&data)

	if err != nil {
		return "", err
	}
	return data, nil
}

func (repo *MongoDBArticleDaoManager) DeleteOne(ctx context.Context, collection string, filter bson.M) error {
	res, err := repo.db.Collection(collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no document deleted")
	}
	return nil
}

func (repo *MongoDBArticleDaoManager) UpdateOne(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (any, error) {
	opts := options.Update().SetUpsert(upsert)
	result, err := repo.db.Collection(collection).UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
	}
	return result, nil
}
