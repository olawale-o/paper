package auth

import (
	"context"
	"go-simple-rest/src/v1/authors"
)

// LoginAuth
// @Description: Request struct for login authentication
type LoginAuth struct {
	// username
	USERNAME string `bson:"username" json:"username" validate:"required,min=1"`
	// password
	PASSWORD string `bson:"username" json:"password" validate:"required,min=4"`
}

type User struct {
	USERNAME string `bson:"username" json:"username"`
	ROLE     string `bson:"role" json:"role"`
	ID       string `bson:"id" json:"id"`
}

type LoginResponse struct {
	MESSAGE string `json:"message"`
	USER    User   `json:"user"`
}

type RegisterAuth struct {
	USERNAME  string `bson:"username" json:"username" validate:"required"`
	PASSWORD  string `bson:"username" json:"password" validate:"required"`
	FIRSTNAME string `bson:"firstname" json:"firstname" validate:"required"`
	LASTNAME  string `bson:"lastname" json:"lastname" validate:"required"`
}

type Repository interface {
	GetUser(ctx context.Context, collection string, username string) (authors.User, error)
	InsertUser(ctx context.Context, collection string, user authors.User) (interface{}, error)
}
