package comment

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
	STATUS          string             `bson:"status,omitempty" json:"status,omitempty"`
	PARENTCOMMENTID primitive.ObjectID `bson:"parentCommentId,omitempty" json:"parentCommentId,omitempty"`
}
