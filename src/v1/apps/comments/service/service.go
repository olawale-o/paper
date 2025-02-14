package service

import (
	"comments/db"
	"comments/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client, ctx, err = db.Connect()

var articleCollection = client.Database("go").Collection("articles")
var collection = client.Database("go").Collection("comments")

func NewComment(c model.Comment, articleId primitive.ObjectID) (error, interface{}) {
	var article model.Article

	filter := bson.M{"_id": articleId}
	if err := articleCollection.FindOne(context.TODO(), filter).Decode(&article); err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return err, "Article not found"
		}
		return err, "Article not found"
	}

	doc := model.Comment{BODY: c.BODY, ARTICLEID: articleId, USERID: c.USERID, LIKES: 0, CREATEDAT: time.Now(), UPDATEDAT: time.Now(), STATUS: "pending", PARENTCOMMENTID: c.PARENTCOMMENTID}
	res, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err)
		return err, ""
	}
	id := res.InsertedID

	return err, id
}

func GetComment(articleId primitive.ObjectID, commentId primitive.ObjectID) (error, interface{}) {
	var comment model.Comment

	filter := bson.M{"_id": commentId, "articleId": articleId}
	if err := collection.FindOne(context.TODO(), filter).Decode(&comment); err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return err, "Comment not found"
		}
		return err, "Comment not found"
	}

	return err, comment
}

func GetComments(articleId primitive.ObjectID) []model.Comment {
	var comments []model.Comment

	filter := bson.M{"articleId": articleId}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	if err = cursor.All(context.TODO(), &comments); err != nil {
		panic(err)
	}
	return comments
}
