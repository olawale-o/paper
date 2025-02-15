package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID              interface{}        `bson:"_id,omitempty" json:"id,omitempty"`
	ARTICLEID       primitive.ObjectID `bson:"articleId" json:"articleId"`
	BODY            string             `bson:"body" json:"body"`
	USERID          primitive.ObjectID `bson:"userId,omitempty" json:"userId,omitempty"`
	LIKES           int                `bson:"likes,omitempty" json:"likes,omitempty"`
	CREATEDAT       time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UPDATEDAT       time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	DELETEDAT       time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
	STATUS          string             `bson:"status,omitempty" json:"status,omitempty"`
	PARENTCOMMENTID primitive.ObjectID `bson:"parentCommentId,omitempty" json:"parentCommentId,omitempty"`
}

type Article struct {
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

type ArticleCommentReq struct {
	BODY string `bson:"body" json:"body"`
}

type ArticleData struct {
	ARTICLEID       string
	BODY            string
	USERID          string
	PARENTCOMMENTID string
}

type Payload struct {
	Data  interface{}
	Event string
}
