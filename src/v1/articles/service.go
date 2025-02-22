package articles

import (
	"context"
	"fmt"
	"go-simple-rest/db"
	"go-simple-rest/src/v1/articles/repo"
	"go-simple-rest/src/v1/authors/model"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client, ctx, err = db.Connect()

var collection = client.Database("go").Collection("articles")

var database = client.Database("go")

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

func GetArticle(c *gin.Context) (interface{}, error) {
	var article bson.M
	r, err := repo.New(database)

	if err != nil {
		return article, err
	}
	oid, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{"_id": oid}
	data, err := r.FindOne(context.TODO(), "articles", filter, article)

	if err != nil {
		return article, err
	}

	return data, nil
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
