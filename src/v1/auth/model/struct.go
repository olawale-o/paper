package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LoginAuth
// @Description: Request struct for login authentication
type LoginAuth struct {
	// username
	USERNAME string `bson:"username" json:"username" validate:"required,min=1"`
	// password
	PASSWORD string `bson:"username" json:"password" validate:"required,min=4"`
}

type UserResponseObject struct {
	USERNAME string `json:"username"`
	ROLE     string `json:"role"`
	ID       string `json:"id"`
}

type LoginResponse struct {
	MESSAGE string             `json:"message"`
	TOKEN   string             `json:"token"`
	USER    UserResponseObject `json:"user"`
}

type RegisterAuth struct {
	USERNAME  string `bson:"username" json:"username" validate:"required"`
	PASSWORD  string `bson:"username" json:"password" validate:"required"`
	FIRSTNAME string `bson:"firstname" json:"firstname" validate:"required"`
	LASTNAME  string `bson:"lastname" json:"lastname" validate:"required"`
}

type User struct {
	ID                string             `bson:"_id,omitempty" json:"id,omitempty"`
	FIRSTNAME         string             `bson:"firstName" json:"firstName"`
	LASTNAME          string             `bson:"lastName" json:"lastName"`
	USERNAME          string             `bson:"username" json:"username"`
	PASSWORD          string             `bson:"password" json:"password"`
	ARTICLECOUNT      int                `bson:"articleCount,omitempty" json:"articleCount,omitempty"`
	ARTICLELIKESCOUNT int                `bson:"articleLikesCount,omitempty" json:"articleLikesCount,omitempty"`
	CREATEDAT         primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt,omitempty" swaggertype:"string"`
	UPDATEDAT         primitive.DateTime `bson:"updatedAt,omitempty" json:"updatedAt,omitempty" swaggertype:"string"`
	ROLE              string             `bson:"role,omitempty" json:"role,omitempty"`
	CREATEDATIMESTAMP int64              `bson:"createdAtTimestamp,omitempty" json:"createdAtTimestamp,omitempty"`
	UPDATEDATIMESTAMP int64              `bson:"updatedAtTimestamp,omitempty" json:"updatedAtTimestamp,omitempty"`
	DELETEDATIMESTAMP int64              `bson:"deletedAtTimestamp,omitempty" json:"deletedAtTimestamp,omitempty"`
}

type Repository interface {
	FindOne(ctx context.Context, collection string, username string) (User, error)
	InsertOne(ctx context.Context, collection string, user User) (interface{}, error)
}
