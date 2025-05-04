package dao

import (
	"context"
	"fmt"
	"log"

	"go-simple-rest/src/v1/authors/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBAuthorDaoManager struct {
	db     *mongo.Database
	logger log.Logger
}

func New(db *mongo.Database) (AuthorDAO, error) {
	return &MongoDBAuthorDaoManager{
		db: db,
	}, nil
}

func (repo *MongoDBAuthorDaoManager) Get(ctx context.Context, collection string, filter bson.M) ([]model.AuthorArticle, error) {
	cursor, err := repo.db.Collection(collection).Find(context.TODO(), filter)

	if err != nil {
		return []model.AuthorArticle{}, err
	}
	var articles []model.AuthorArticle
	if err = cursor.All(context.TODO(), &articles); err != nil {
		return []model.AuthorArticle{}, err
	}
	return articles, nil
}

func (repo *MongoDBAuthorDaoManager) FindOne(ctx context.Context, collection string, filter bson.M, v model.Author) (model.Author, error) {
	if err := repo.db.Collection(collection).FindOne(context.TODO(), filter).Decode(&v); err != nil {
		return v, err
	}
	return v, nil
}

func (repo *MongoDBAuthorDaoManager) InsertOne(ctx context.Context, collection string, doc any) (any, error) {
	res, err := repo.db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		return "", err
	}
	return res.InsertedID, nil
}

func (repo *MongoDBAuthorDaoManager) FindOneAndUpdate(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (model.AuthorArticleUpdateResponse, error) {
	var data model.AuthorArticleUpdateResponse
	opts := options.FindOneAndUpdate().SetUpsert(upsert)
	repo.db.Collection(collection).FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&data)

	return data, nil
}

func (repo *MongoDBAuthorDaoManager) DeleteOne(ctx context.Context, collection string, filter bson.M) error {
	res, err := repo.db.Collection(collection).DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no document deleted")
	}
	return nil
}

func (repo *MongoDBAuthorDaoManager) UpdateOne(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (any, error) {
	opts := options.Update().SetUpsert(upsert)
	result, err := repo.db.Collection(collection).UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
	}
	return result, nil
}
