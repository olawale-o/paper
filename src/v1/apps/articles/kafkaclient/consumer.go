package kafkaclient

import (
	"articles/db"
	"articles/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/IBM/sarama"
	"go.mongodb.org/mongo-driver/bson"
)

var client, ctx, err = db.Connect()

var collection = client.Database("go").Collection("article_views")

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

		if msg.Topic == "article.views" {
			var article model.ArticleView
			err := json.Unmarshal(msg.Value, &article)
			if err != nil {
				log.Printf("unable to unmarshal article view: %v", err)
				continue
			} else {
				SaveArticleView(article)
			}
		}

		sess.MarkMessage(msg, "")
	}
	return nil
}

func SaveArticleView(article model.ArticleView) error {
	var data interface{}
	filter := bson.M{"userId": article.USERID, "articleId": article.ID}
	update := bson.M{"userId": article.USERID, "articleId": article.ID, "createdAtTimestamp": article.CREATEDATTIMESTAMP}
	collection.FindOneAndUpdate(ctx, filter, update).Decode(&data)
	return nil
}
