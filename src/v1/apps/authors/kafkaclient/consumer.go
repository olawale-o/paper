package kafkaclient

import (
	"authors/db"
	"authors/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/IBM/sarama"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client, ctx, err = db.Connect()

var collection = client.Database("go").Collection("users")

func ConsumerWithAutoCommit(groupName string) {
	config := KafkaConfig()
	client, err := sarama.NewConsumerGroup(BROKERS, groupName, config)
	if err != nil {
		log.Fatalf("unable to create kafka consumer group: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithCancel(context.Background())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			err := client.Consume(ctx, TOPICS, &ConsumerHandler{})
			if err != nil {
				log.Printf("consume error: %v", err)
			}

			select {
			case <-signals:
				cancel()
				return
			default:
			}
		}
	}()

	wg.Wait()
}

type ConsumerHandler struct{}

func (h *ConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ConsumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Received message: key=%s, value=%s, partition=%d, offset=%d\n", string(msg.Key), string(msg.Value), msg.Partition, msg.Offset)

		if msg.Topic == "articles.new" {
			var article model.ArticleData
			err := json.Unmarshal(msg.Value, &article)
			if err != nil {
				log.Printf("unable to unmarshal article data: %v", err)
				continue
			} else {
				UpdateAuthor(article)
			}
		}

		sess.MarkMessage(msg, "")
	}
	return nil
}

func UpdateAuthor(article model.ArticleData) error {
	var data interface{}
	filter := bson.M{"_id": article.AUTHORID}
	update := bson.M{
		"$push": bson.M{"articles": bson.M{"$each": []model.Article{{
			TITLE:              article.TITLE,
			ID:                 article.ID,
			CONTENT:            article.CONTENT,
			CATEGORIES:         article.CATEGORIES,
			TAGS:               article.TAGS,
			CREATEDATTIMESTAMP: article.CREATEDATTIMESTAMP,
			UPDATEDATTIMESTAMP: article.UPDATEDATTIMESTAMP,
		}},
			"$sort":  bson.M{"createdAtTimestamp": -1},
			"$slice": 2}},
		"$inc": bson.M{"articleCount": 1}}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&data)
	return nil
}
