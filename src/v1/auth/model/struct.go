package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LoginAuth
// @Description: Request struct for login authentication
type LoginAuth struct {
	// username
	USERNAME string `bson:"username" json:"username" validate:"required,min=1"`
	// password
	PASSWORD string `bson:"password" json:"password" validate:"required,min=4"`
}

type UserResponseObject struct {
	USERNAME string `json:"username"`
	ROLE     string `json:"role"`
	ID       string `json:"id"`
}

type AuthResponse struct {
	MESSAGE string      `json:"message"`
	DATA    interface{} `json:"data"`
}

type RegisterAuth struct {
	USERNAME  string `bson:"username" json:"username" validate:"required"`
	PASSWORD  string `bson:"password" json:"password" validate:"required"`
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
