package repo

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

func (repo *repository) Get(ctx context.Context, collection string, filter bson.M, sort bson.M, opts bson.M) ([]model.Article, error) {
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

func (repo *repository) FindOne(ctx context.Context, collection string, filter bson.M, v bson.M, opts bson.M) (interface{}, error) {
	options := options.FindOne().SetProjection(opts)
	if err := repo.db.Collection(collection).FindOne(context.TODO(), filter, options).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

func (repo *repository) InsertOne(ctx context.Context, collection string, doc interface{}) (interface{}, error) {
	res, err := repo.db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		return "", err
	}
	return res.InsertedID, nil
}

func (repo *repository) FindOneAndUpdate(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (interface{}, error) {
	var data interface{}
	opts := options.FindOneAndUpdate().SetUpsert(upsert)
	repo.db.Collection(collection).FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&data)

	if err != nil {
		return "", err
	}
	return data, nil
}

func (repo *repository) DeleteOne(ctx context.Context, collection string, filter bson.M) error {
	res, err := repo.db.Collection(collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no document deleted")
	}
	return nil
}

func (repo *repository) UpdateOne(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (interface{}, error) {
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
