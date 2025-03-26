package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleInteraction struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ARTICLEID         primitive.ObjectID `bson:"articleId,omitempty" json:"articleId,omitempty"`
	USERID            primitive.ObjectID `bson:"userId,omitempty" json:"userId,omitempty"`
	TYPE              string             `bson:"type,omitempty" json:"type,omitempty"`
	CREATEDAT         primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt,omitempty" swaggertype:"string"`
	CREATEDATIMESTAMP int64              `bson:"createdAtTimestamp,omitempty" json:"createdAtTimestamp,omitempty"`
}

type Article struct {
	ID                interface{}        `bson:"_id,omitempty" json:"id,omitempty"`
	TITLE             string             `bson:"title" json:"title"`
	CONTENT           string             `bson:"content" json:"content"`
	AUTHORID          primitive.ObjectID `bson:"authorId,omitempty" json:"authorId,omitempty"`
	LIKES             int                `bson:"likes,omitempty" json:"likes,omitempty"`
	VIEWS             int                `bson:"views,omitempty" json:"views,omitempty"`
	CREATEDAT         primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt,omitempty" swaggertype:"string"`
	UPDATEDAT         primitive.DateTime `bson:"updatedAt,omitempty" json:"updatedAt,omitempty" swaggertype:"string"`
	STATUS            string             `bson:"status,omitempty" json:"status,omitempty"`
	CATEGORIES        []string           `bson:"categories,omitempty" json:"categories,omitempty"`
	TAGS              []string           `bson:"tags,omitempty" json:"tags,omitempty"`
	DELETEDAT         primitive.DateTime `bson:"deletedAt,omitempty" json:"deletedAt,omitempty" swaggertype:"string"`
	CREATEDATIMESTAMP int                `bson:"createdAtTimestamp,omitempty" json:"createdAtTimestamp,omitempty"`
	UPDATEDATIMESTAMP int                `bson:"updatedAtTimestamp,omitempty" json:"updatedAtTimestamp,omitempty"`
	DELETEDATIMESTAMP int                `bson:"deletedAtTimestamp,omitempty" json:"deletedAtTimestamp,omitempty"`
}

type Author struct {
	ID        interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	FIRSTNAME string      `bson:"firstName" json:"firstName"`
	LASTNAME  string      `bson:"lastName" json:"lastName"`
	USERNAME  string      `bson:"username" json:"username"`
}

type AuthorArticle struct {
	ID         interface{}        `bson:"_id,omitempty" json:"id,omitempty"`
	TITLE      string             `bson:"title" json:"title"`
	CONTENT    string             `bson:"content" json:"content"`
	AUTHOR     Author             `bson:"authorId,omitempty" json:"authorId,omitempty"`
	LIKES      int                `bson:"likes,omitempty" json:"likes,omitempty"`
	VIEWS      int                `bson:"views,omitempty" json:"views,omitempty"`
	CREATEDAT  primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UPDATEDAT  primitive.DateTime `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	STATUS     string             `bson:"status,omitempty" json:"status,omitempty"`
	CATEGORIES []string           `bson:"categories,omitempty" json:"categories,omitempty"`
	TAGS       []string           `bson:"tags,omitempty" json:"tags,omitempty"`
	DELETEDAT  primitive.DateTime `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}

type QueryParams struct {
	Date  string `json:"date"`
	Likes string `json:"likes"`
	Views string `json:"views"`
}
