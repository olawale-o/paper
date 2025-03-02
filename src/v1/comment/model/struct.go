package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID                interface{}        `bson:"_id,omitempty" json:"id,omitempty"`
	ARTICLEID         primitive.ObjectID `bson:"articleId" json:"articleId"`
	BODY              string             `bson:"body" json:"body"`
	USERID            primitive.ObjectID `bson:"userId,omitempty" json:"userId,omitempty"`
	LIKES             int                `bson:"likes,omitempty" json:"likes,omitempty"`
	CREATEDAT         primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UPDATEDAT         primitive.DateTime `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	DELETEDAT         primitive.DateTime `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
	STATUS            string             `bson:"status,omitempty" json:"status,omitempty"`
	PARENTCOMMENTID   primitive.ObjectID `bson:"parentCommentId,omitempty" json:"parentCommentId,omitempty"`
	CREATEDATIMESTAMP int64              `bson:"createdAtTimestamp,omitempty" json:"createdAtTimestamp,omitempty"`
	UPDATEDATIMESTAMP int64              `bson:"updatedAtTimestamp,omitempty" json:"updatedAtTimestamp,omitempty"`
	DELETEDATIMESTAMP int64              `bson:"deletedAtTimestamp,omitempty" json:"deletedAtTimestamp,omitempty"`
}

type CommentArticle struct {
	ID                interface{}        `bson:"_id,omitempty" json:"id,omitempty"`
	TITLE             string             `bson:"title" json:"title"`
	CONTENT           string             `bson:"content" json:"content"`
	AUTHORID          primitive.ObjectID `bson:"authorId,omitempty" json:"authorId,omitempty"`
	LIKES             int                `bson:"likes,omitempty" json:"likes,omitempty"`
	VIEWS             int                `bson:"views,omitempty" json:"views,omitempty"`
	CREATEDAT         primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UPDATEDAT         primitive.DateTime `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	STATUS            string             `bson:"status,omitempty" json:"status,omitempty"`
	CATEGORIES        []string           `bson:"categories,omitempty" json:"categories,omitempty"`
	TAGS              []string           `bson:"tags,omitempty" json:"tags,omitempty"`
	DELETEDAT         time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
	CREATEDATIMESTAMP int                `bson:"createdAtTimestamp,omitempty" json:"createdAtTimestamp,omitempty"`
	UPDATEDATIMESTAMP int                `bson:"updatedAtTimestamp,omitempty" json:"updatedAtTimestamp,omitempty"`
	DELETEDATIMESTAMP int                `bson:"deletedAtTimestamp,omitempty" json:"deletedAtTimestamp,omitempty"`
}

type Repository interface {
	Get(ctx context.Context, collection string, filter bson.M, sort bson.M, limit int64) ([]Comment, error)
	FindOne(ctx context.Context, collection string, filter bson.M, v bson.M) (interface{}, error)
}
