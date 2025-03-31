package repo

import (
	"context"
	"go-simple-rest/src/v1/auth/model"
)

type Repository interface {
	FindOne(ctx context.Context, collection string, username string) (model.User, error)
	InsertOne(ctx context.Context, collection string, user model.User) (interface{}, error)
}
