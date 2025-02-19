package authors

import (
	"go-simple-rest/src/v1/articles"
	"time"
)

type User struct {
	ID                string             `bson:"_id,omitempty" json:"id,omitempty"`
	FIRSTNAME         string             `bson:"firstName" json:"firstName"`
	LASTNAME          string             `bson:"lastName" json:"lastName"`
	USERNAME          string             `bson:"username" json:"username"`
	PASSWORD          string             `bson:"password" json:"password"`
	ARTICLECOUNT      int                `bson:"articleCount,omitempty" json:"articleCount,omitempty"`
	ARTICLELIKESCOUNT int                `bson:"articleLikesCount,omitempty" json:"articleLikesCount,omitempty"`
	CREATEDAT         string             `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	ARTICLES          []articles.Article `bson:"articles,omitempty" json:"articles,omitempty"`
	ROLE              string             `bson:"role,omitempty" json:"role,omitempty"`
}


type Author struct {
	ID        interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	FIRSTNAME string      `bson:"firstName" json:"firstName"`
	LASTNAME  string      `bson:"lastName" json:"lastName"`
	USERNAME  string      `bson:"username" json:"username"`
	CREATEDAT time.Time   `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UPDATEDAT time.Time   `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
