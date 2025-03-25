package model

import (
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

type Reply struct {
	ID                 interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	BODY               string      `bson:"body,omitempty" json:"body,omitempty"`
	COMMENTID          interface{} `bson:"commentId,omitempty" json:"commentId,omitempty"`
	CREATEDATTIMESTAMP int         `bson:"createdAtTimestamp,omitempty" json:"createdAtTimestamp,omitempty"`
}

type ArticleWithComments struct {
	ID                interface{} `bson:"id,omitempty" json:"id,omitempty"`
	ARTICLEID         interface{} `bson:"articleId,omitempty" json:"articleId,omitempty"`
	BODY              string      `bson:"body,omitempty" json:"body,omitempty"`
	REPLIES           []Reply     `bson:"replies,omitempty" json:"replies,omitempty"`
	CREATEDATIMESTAMP int         `bson:"createdAtTimestamp,omitempty" json:"createdAtTimestamp,omitempty"`
}
