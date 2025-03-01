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
	USERNAME string `bson:"username" json:"username"`
	ROLE     string `bson:"role" json:"role"`
	ID       string `bson:"id" json:"id"`
}

type LoginResponse struct {
	MESSAGE            string `json:"message"`
	UserResponseObject User   `json:"user"`
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
	GetUser(ctx context.Context, collection string, username string) (User, error)
	InsertUser(ctx context.Context, collection string, user User) (interface{}, error)
}
