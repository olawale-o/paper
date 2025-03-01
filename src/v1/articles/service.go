package articles

import (
	"context"
	"go-simple-rest/db"

	"go-simple-rest/src/v1/articles/model"
	"go-simple-rest/src/v1/articles/repo"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client, ctx, err = db.Connect()

var collection = client.Database("go").Collection("articles")

var database = client.Database("go")

func handleQueryParams(date, likes, views string) bson.M {
	fieldValues := map[string]int{
		"asc":  1,
		"desc": -1,
	}
	filter := bson.M{}
	if date == "desc" {
		filter["createdAtTimestamp"] = -1
	} else {
		filter["createdAtTimestamp"] = fieldValues[date]
	}
	// if likes == "desc" {
	// 	filter["likes"] = -1
	// } else {
	// 	filter["likes"] = fieldValues[likes]
	// }
	// if views == "desc" {
	// 	filter["views"] = -1
	// } else {
	// 	filter["views"] = fieldValues[views]
	// }
	return filter
}

func GetAll(date, likes, views string) (interface{}, error) {
	filter := bson.M{}
	sort := handleQueryParams(date, likes, views)
	var fields bson.M = bson.M{"deletedAt": 0, "tags": 0, "categories": 0}
	r, err := repo.New(database)

	if err != nil {
		return nil, err
	}
	articles, err := r.Get(context.TODO(), "articles", filter, sort, fields)
	return articles, nil
}

func GetArticle(c *gin.Context) (interface{}, error) {
	var fields bson.M = bson.M{"deletedAt": 0, "tags": 0, "categories": 0}
	var article bson.M
	r, err := repo.New(database)

	if err != nil {
		return article, err
	}
	oid, _ := primitive.ObjectIDFromHex(c.Param("id"))

	filter := bson.M{"_id": oid}
	data, err := r.FindOne(context.TODO(), "articles", filter, article, fields)

	if err != nil {
		return article, err
	}

	return data, nil
}

func Update(c *gin.Context) (interface{}, error) {
	r, err := repo.New(database)

	if err != nil {
		return nil, err
	}
	oid, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var updatedArticle model.Article
	if err := c.BindJSON(&updatedArticle); err != nil {
		log.Println(err)
		return updatedArticle, err
	}
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": bson.M{"title": updatedArticle.TITLE, "author": updatedArticle.AUTHORID}}

	result, err := r.UpdateOne(context.TODO(), "articles", filter, update, true)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return updatedArticle, err
		}
	}

	return result, nil
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
