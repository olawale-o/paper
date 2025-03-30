package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginAuth struct {
	USERNAME string `bson:"username" json:"username" validate:"required,min=1"`
	PASSWORD string `bson:"username" json:"password" validate:"required,min=4"`
}

type RegisterAuth struct {
	USERNAME  string `bson:"username" json:"username" validate:"required"`
	PASSWORD  string `bson:"username" json:"password" validate:"required"`
	FIRSTNAME string `bson:"firstname" json:"firstname" validate:"required"`
	LASTNAME  string `bson:"lastname" json:"lastname" validate:"required"`
}

type Article struct {
	ID         interface{}        `bson:"_id,omitempty" json:"id,omitempty"`
	TITLE      string             `bson:"title" json:"title"`
	CONTENT    string             `bson:"content" json:"content"`
	AUTHORID   primitive.ObjectID `bson:"authorId,omitempty" json:"authorId,omitempty"`
	LIKES      int                `bson:"likes,omitempty" json:"likes,omitempty"`
	VIEWS      int                `bson:"views,omitempty" json:"views,omitempty"`
	CREATEDAT  primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt,omitempty" swaggertype:"string"`
	UPDATEDAT  primitive.DateTime `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	STATUS     string             `bson:"status,omitempty" json:"status,omitempty"`
	CATEGORIES []string           `bson:"categories,omitempty" json:"categories,omitempty"`
	TAGS       []string           `bson:"tags,omitempty" json:"tags,omitempty"`
	DELETEDAT  primitive.DateTime `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}

type ArticleCommentReq struct {
	BODY string `bson:"body" json:"body"`
}

type User struct {
	ID                string             `bson:"_id,omitempty" json:"id,omitempty"`
	FIRSTNAME         string             `bson:"firstName" json:"firstName"`
	LASTNAME          string             `bson:"lastName" json:"lastName"`
	USERNAME          string             `bson:"username" json:"username"`
	PASSWORD          string             `bson:"password" json:"password"`
	ARTICLECOUNT      int                `bson:"articleCount,omitempty" json:"articleCount,omitempty"`
	ARTICLELIKESCOUNT int                `bson:"articleLikesCount,omitempty" json:"articleLikesCount,omitempty"`
	CREATEDAT         primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	ARTICLES          []Article          `bson:"articles,omitempty" json:"articles,omitempty"`
	ROLE              string             `bson:"role,omitempty" json:"role,omitempty"`
}

type UserResponseObject struct {
	USERNAME string `json:"username"`
	ROLE     string `json:"role"`
	ID       string `json:"id"`
}

type LoginResponse struct {
	// MESSAGE string             `json:"message"`
	TOKEN string             `json:"token"`
	USER  UserResponseObject `json:"user"`
}

type Repository interface {
	GetUser(ctx context.Context, collection string, username string) (User, error)
	InsertUser(ctx context.Context, collection string, user User) (interface{}, error)
}
