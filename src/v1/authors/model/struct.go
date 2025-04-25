package model

import (
	"time"

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
	TITLE      string             `bson:"title" json:"title" validate:"required,min=1"`
	CONTENT    string             `bson:"content" json:"content" validate:"required,min=1"`
	AUTHORID   primitive.ObjectID `bson:"authorId,omitempty" json:"authorId,omitempty"`
	LIKES      int                `bson:"likes,omitempty" json:"likes,omitempty"`
	VIEWS      int                `bson:"views,omitempty" json:"views,omitempty"`
	CREATEDAT  time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UPDATEDAT  time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	STATUS     string             `bson:"status,omitempty" json:"status,omitempty"`
	CATEGORIES []string           `bson:"categories,omitempty" json:"categories,omitempty" validate:"required"`
	TAGS       []string           `bson:"tags,omitempty" json:"tags,omitempty" validate:"required"`
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
	ID interface{} `bson:"_id,omitempty" json:"id,omitempty"`
}
