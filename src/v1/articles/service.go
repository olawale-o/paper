package articles

import (
	"context"
	"fmt"
	"go-simple-rest/db"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client, ctx, err = db.Connect()

var collection = client.Database("go").Collection("articles")

func GetAll() []Article {
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	var articles []Article
	if err = cursor.All(context.TODO(), &articles); err != nil {
		panic(err)
	}
	return articles
}

func CreateArticle(c *gin.Context) (error, interface{}) {
	var newArticle Article
	if err := c.BindJSON(&newArticle); err != nil {
		log.Println(err)
		return err, ""
	}
	doc := Article{TITLE: newArticle.TITLE, AUTHOR: newArticle.AUTHOR}
	res, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err)
		return err, ""
	}
	id := res.InsertedID

	return err, id
}

func GetArticle(c *gin.Context) (Article, error) {
	oid, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var article Article
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
	var updatedArticle Article
	if err := c.BindJSON(&updatedArticle); err != nil {
		log.Println(err)
		return updatedArticle, err
	}
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": bson.M{"title": updatedArticle.TITLE, "author": updatedArticle.AUTHOR}}
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
