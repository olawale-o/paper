package repo

import (
	"context"
	"fmt"
	"go-simple-rest/db"
	"log"

	"go-simple-rest/src/v1/authors/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (repo *repository) Get(ctx context.Context, collection string, filter bson.M) ([]model.AuthorArticle, error) {
	cursor, err := repo.db.Collection(collection).Find(context.TODO(), filter)

	if err != nil {
		fmt.Println(err.Error())
		return []model.AuthorArticle{}, err
	}
	var articles []model.AuthorArticle
	if err = cursor.All(context.TODO(), &articles); err != nil {
		return []model.AuthorArticle{}, err
	}
	return articles, nil
}

func (repo *repository) InsertUser(ctx context.Context, collection string, user model.User) (interface{}, error) {
	doc := model.User{
		FIRSTNAME: user.FIRSTNAME,
		LASTNAME:  user.LASTNAME,
		USERNAME:  user.USERNAME,
		PASSWORD:  user.PASSWORD,
		ROLE:      "author",
		CREATEDAT: time.Now().Format(time.DateTime),
	}
	res, err := repo.db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		return "", err
	}
	return res.InsertedID, nil
}
