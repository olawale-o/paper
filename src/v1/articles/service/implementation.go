package service

import (
	"context"
	"encoding/json"
	"go-simple-rest/db"
	"time"

	"go-simple-rest/src/v1/articles/model"
	"go-simple-rest/src/v1/articles/repository"
	"go-simple-rest/src/v1/kafkaclient"
	"go-simple-rest/src/v1/utils"
	"log"

	"github.com/IBM/sarama"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client, ctx, err = db.Connect()

type ServiceManager struct {
	repo repository.Repository
}

func New(repo repository.Repository) (Service, error) {
	return &ServiceManager{repo: repo}, nil
}

func (sm *ServiceManager) GetAll(params model.QueryParams) (interface{}, error) {
	filter := bson.M{}
	sort := utils.HandleQueryParams(params)
	var fields bson.M = bson.M{"deletedAt": 0, "tags": 0, "categories": 0}
	articles, err := sm.repo.Get(context.TODO(), "articles", filter, sort, fields)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (sm *ServiceManager) GetArticle(articleId primitive.ObjectID) (interface{}, error) {
	var fields bson.M = bson.M{"deletedAt": 0, "tags": 0, "categories": 0}
	var article bson.M

	filter := bson.M{"_id": articleId}
	data, err := sm.repo.FindOne(context.TODO(), "articles", filter, article, fields)

	if err != nil {
		return article, err
	}

	_, err = json.Marshal(model.ArticleInteraction{
		ARTICLEID:         articleId,
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

func (sm *ServiceManager) Update(articleId primitive.ObjectID, article model.Article) (interface{}, error) {

	filter := bson.M{"_id": articleId}
	update := bson.M{"$set": bson.M{"title": article.TITLE, "author": article.AUTHORID}}

	result, err := sm.repo.UpdateOne(context.TODO(), "articles", filter, update, true)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return result, err
		}
	}

	return result, nil
}

// func (sm *ServiceManager) Delete(c *gin.Context) (int64, error) {
// oid, _ := primitive.ObjectIDFromHex(c.Param("id"))
// filter := bson.M{"_id": oid}
// result, err := collection.DeleteOne(context.TODO(), filter)
// if err != nil {
// 	log.Println(err)
// 	return 0, err
// }

// return result.DeletedCount, nil
// return 0, nil
// }
