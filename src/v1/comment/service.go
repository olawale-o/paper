package comment

import (
	"context"
	"go-simple-rest/db"
	"go-simple-rest/src/v1/articles"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client, ctx, err = db.Connect()

var articleCollection = client.Database("go").Collection("articles")
var collection = client.Database("go").Collection("comments")

func NewComment(c Comment, articleId primitive.ObjectID) (error, interface{}) {
	var article articles.Article

	filter := bson.M{"_id": articleId}
	if err := articleCollection.FindOne(context.TODO(), filter).Decode(&article); err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return err, "Article not found"
		}
		return err, "Article not found"
	}

	doc := Comment{BODY: c.BODY, ARTICLEID: articleId, USERID: c.USERID, LIKES: 0, CREATEDAT: time.Now(), UPDATEDAT: time.Now(), STATUS: "pending", PARENTCOMMENTID: c.PARENTCOMMENTID}
	res, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err)
		return err, ""
	}
	id := res.InsertedID

	return err, id
}
