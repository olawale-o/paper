package articles

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID        interface{}        `bson:"_id,omitempty" json:"id,omitempty"`
	TITLE     string             `bson:"title" json:"title"`
	CONTENT   string             `bson:"content" json:"content"`
	AUTHORID  primitive.ObjectID `bson:"authorId,omitempty" json:"authorId,omitempty"`
	LIKES     int                `bson:"likes,omitempty" json:"likes,omitempty"`
	VIEWS     int                `bson:"views,omitempty" json:"views,omitempty"`
	CREATEDAT time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UPDATEDAT time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}
