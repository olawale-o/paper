package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Author struct {
	ID        interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	FIRSTNAME string      `bson:"firstName" json:"firstName"`
	LASTNAME  string      `bson:"lastName" json:"lastName"`
	USERNAME  string      `bson:"username" json:"username"`
	CREATEDAT time.Time   `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UPDATEDAT time.Time   `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type AuthorArticle struct {
	ID         interface{}        `bson:"_id,omitempty" json:"id,omitempty"`
	TITLE      string             `bson:"title" json:"title"`
	CONTENT    string             `bson:"content" json:"content"`
	AUTHORID   primitive.ObjectID `bson:"authorId,omitempty" json:"authorId,omitempty"`
	LIKES      int                `bson:"likes,omitempty" json:"likes,omitempty"`
	VIEWS      int                `bson:"views,omitempty" json:"views,omitempty"`
	CREATEDAT  time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UPDATEDAT  time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	STATUS     string             `bson:"status,omitempty" json:"status,omitempty"`
	CATEGORIES []string           `bson:"categories,omitempty" json:"categories,omitempty"`
	TAGS       []string           `bson:"tags,omitempty" json:"tags,omitempty"`
	DELETEDAT  time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}

type AuthorArticleUpdateRequest struct {
	TITLE   string `bson:"title" json:"title"`
	CONTENT string `bson:"content" json:"content"`
}

type AuthorArticleResponse struct {
	MESSAGE  string          `json:"message,omitempty"`
	ARTICLES []AuthorArticle `json:"articles,omitempty"`
}

type AuthorArticleUpdateResponse struct {
	MESSAGE string `json:"message,omitempty"`
}

type Repository interface {
	Get(ctx context.Context, collection string, filter bson.M) ([]AuthorArticle, error)
	FindOne(ctx context.Context, collection string, filter bson.M, v bson.M) (interface{}, error)
	InsertOne(ctx context.Context, collection string, doc interface{}) (interface{}, error)
	FindOneAndUpdate(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (interface{}, error)
	DeleteOne(ctx context.Context, collection string, filter bson.M) error
	UpdateOne(ctx context.Context, collection string, filter bson.M, update bson.M, upsert bool) (interface{}, error)
}
