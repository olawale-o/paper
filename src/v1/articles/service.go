package articles

import (
	"context"
	"encoding/json"
	"go-simple-rest/db"
	"time"

	"go-simple-rest/src/v1/articles/model"
	"go-simple-rest/src/v1/articles/repo"
	"go-simple-rest/src/v1/kafkaclient"
	"go-simple-rest/src/v1/utils"
	"log"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client, ctx, err = db.Connect()

var collection = client.Database("go").Collection("articles")

var database = client.Database("go")

func GetAll(params model.QueryParams) (interface{}, error) {
	filter := bson.M{}
	sort := utils.HandleQueryParams(params)
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
	oid, err := utils.ParseParamToPrimitiveObjectId(c.Param("id"))

	if err != nil {
		return article, err
	}

	filter := bson.M{"_id": oid}
	data, err := r.FindOne(context.TODO(), "articles", filter, article, fields)

	if err != nil {
		return article, err
	}

	_, err = json.Marshal(model.ArticleInteraction{
		ARTICLEID:         oid,
		TYPE:              "view",
		CREATEDAT:         primitive.NewDateTimeFromTime(time.Now()),
		CREATEDATIMESTAMP: time.Now().Local().UnixMilli(),
	})
	if err != nil {
		return article, err
	}
	producer := kafkaclient.KafkaAsyncProducer()
	message := &sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder("Hello World! async producer")}

	kafkaclient.ProduceAsyncMessage(producer, message)

	// natsclient.PublishMessage(context.Background(), "INTERACTIONS.view", value)

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
