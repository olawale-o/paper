package repo

import (
	"auth/db"
	"auth/model"
	"context"

	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client, ctx, err = db.Connect()

var collection = client.Database("go").Collection("users")

type repository struct {
	db     *mongo.Database
	logger log.Logger
}

func New(db *mongo.Database) (model.Repository, error) {
	// return  repository
	return &repository{
		db: db,
	}, nil
}

func (repo *repository) GetUser(ctx context.Context, collection string, username string) (model.User, error) {
	var dbUser model.User
	err := repo.db.Collection(collection).FindOne(context.TODO(), bson.M{"username": username}).Decode(&dbUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return dbUser, err
		}
		return dbUser, err
	}
	return dbUser, nil
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
