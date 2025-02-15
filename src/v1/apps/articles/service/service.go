package service

import (
	"articles/db"
	"articles/events"
	"articles/model"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client, ctx, err = db.Connect()

var collection = client.Database("go").Collection("articles")

func GetAll() []model.Article {
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	var articles []model.Article
	if err = cursor.All(context.TODO(), &articles); err != nil {
		panic(err)
	}
	return articles
}

func CreateArticle(c *gin.Context) (error, interface{}) {
	var newArticle model.Article
	if err := c.BindJSON(&newArticle); err != nil {
		log.Println(err)
		return err, ""
	}

	doc := model.Article{TITLE: newArticle.TITLE, AUTHORID: newArticle.AUTHORID, CONTENT: newArticle.CONTENT}
	res, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err)
		return err, ""
	}
	id := res.InsertedID

	return err, id
}

func GetArticle(c *gin.Context) (model.Article, error) {
	oid, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var article model.Article
	filter := bson.M{"_id": oid}
	if err := collection.FindOne(context.TODO(), filter).Decode(&article); err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return article, err
		}
		return article, err
	}
	return article, nil
}

func Update(c *gin.Context) (interface{}, error) {
	oid, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var updatedArticle model.Article
	if err := c.BindJSON(&updatedArticle); err != nil {
		log.Println(err)
		return updatedArticle, err
	}
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": bson.M{"title": updatedArticle.TITLE, "author": updatedArticle.AUTHORID}}
	opts := options.Update().SetUpsert(true)

	result, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return updatedArticle, err
		}
	}

	return result.UpsertedID, nil
}

func Delete(c *gin.Context) (int64, error) {
	oid, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.M{"_id": oid}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return result.DeletedCount, nil
}

// func GetComments(articleId string) []model.Comment {
// 	filter := bson.M{"articleid": articleId}
// 	var comments []model.Comment
// 	// produce event to comment endpoint
// 	return comments
// }

func CreateComment(articleId string, comment model.ArticleComment) error {
	id, _ := primitive.ObjectIDFromHex(articleId)

	filter := bson.M{"_id": id}
	var article model.Article
	if err = collection.FindOne(context.TODO(), filter).Decode(&article); err != nil {
		if err == mongo.ErrNoDocuments {
			return err
		}
		return err
	}
	events.PublishCommentEvent(
		model.Payload{
			Event: "NEW_COMMENT",
			Data:  model.CommentData{ARTICLEID: articleId, USERID: comment.USERID, BODY: comment.BODY, PARENTCOMMENTID: comment.PARENTCOMMENTID},
		})
	return nil
}
