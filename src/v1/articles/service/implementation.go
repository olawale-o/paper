package service

import (
	"encoding/json"
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

type ServiceManager struct {
	articleRepo repository.Repostiory
}

func New(articleRepo repository.Repostiory) (Service, error) {
	return &ServiceManager{articleRepo: articleRepo}, nil
}

func (sm *ServiceManager) GetAll(params model.QueryParams) ([]model.Article, error) {
	filter := bson.M{}
	sort := utils.HandleQueryParams(params)

	articles, err := sm.articleRepo.GetArticles(filter, sort)
	if err != nil {
		return []model.Article{}, err
	}
	return articles, nil
}

func (sm *ServiceManager) GetArticle(articleId primitive.ObjectID) (any, error) {

	var article model.Article
	data, err := sm.articleRepo.GetArticleById(articleId)

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

	// if err != nil {
	// 	return article, err
	// }

	return data, nil
}

func (sm *ServiceManager) Update(articleId primitive.ObjectID, article model.Article) (any, error) {

	filter := bson.M{"_id": articleId}
	update := bson.M{"$set": bson.M{"title": article.TITLE, "author": article.AUTHORID}}

	result, err := sm.articleRepo.UpdateArticle(filter, update)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return result, err
		}
	}

	return result, nil
}
