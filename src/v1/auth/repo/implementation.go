package repo

import (
	"context"
	"log"

	"go-simple-rest/src/v1/auth/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	db     *mongo.Database
	logger *log.Logger
}

func New(db *mongo.Database) (*AuthRepository, error) {
	return &AuthRepository{
		db: db,
	}, nil
}

func (repo *AuthRepository) FindOne(ctx context.Context, collection string, username string) (model.User, error) {
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

func (repo *AuthRepository) InsertOne(ctx context.Context, collection string, user model.User) (any, error) {
	doc := model.User{
		FIRSTNAME:         user.FIRSTNAME,
		LASTNAME:          user.LASTNAME,
		USERNAME:          user.USERNAME,
		PASSWORD:          user.PASSWORD,
		ROLE:              "author",
		CREATEDAT:         primitive.NewDateTimeFromTime(time.Now()),
		UPDATEDAT:         primitive.NewDateTimeFromTime(time.Now()),
		CREATEDATIMESTAMP: time.Now().Local().UnixMilli(),
		UPDATEDATIMESTAMP: time.Now().Local().UnixMilli(),
	}
	res, err := repo.db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		return "", err
	}
	return res.InsertedID, nil
}
