package natsclient

import (
	"context"
	"encoding/json"
	"fmt"
	"go-simple-rest/db"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleInteraction struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ARTICLEID         primitive.ObjectID `bson:"articleId,omitempty" json:"articleId,omitempty"`
	USERID            primitive.ObjectID `bson:"userId,omitempty" json:"userId,omitempty"`
	TYPE              string             `bson:"type,omitempty" json:"type,omitempty"`
	CREATEDAT         primitive.DateTime `bson:"createdAt,omitempty" json:"createdAt,omitempty" swaggertype:"string"`
	CREATEDATIMESTAMP int64              `bson:"createdAtTimestamp,omitempty" json:"createdAtTimestamp,omitempty"`
}

var client, ctx, err = db.Connect()

var collection = client.Database("go").Collection("article_interactions")

func natsConnect() *nats.Conn {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		fmt.Println("a")
		log.Println(err)
	}

	return nc
}

func jetstreamConnect() jetstream.JetStream {
	js, err := jetstream.New(natsConnect())

	if err != nil {
		log.Println(err)
	}

	return js
}

func getStream() jetstream.Stream {
	js := jetstreamConnect()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stream, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "INTERACTIONS",
		Subjects: []string{"INTERACTIONS.*"},
	})

	if err != nil {
		log.Println(err)
	}

	return stream
}

func getConsumer(stream jetstream.Stream) jetstream.Consumer {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	consumer, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "NEW",
		Name:      "NEW",
		AckPolicy: jetstream.AckExplicitPolicy,
	})

	if err != nil {
		log.Println(err)
	}

	return consumer
}

func Consumer() (jetstream.ConsumeContext, error) {
	cons := getConsumer(getStream())
	cc, err := cons.Consume(func(msg jetstream.Msg) {
		var val ArticleInteraction
		err := json.Unmarshal(msg.Data(), &val)
		if err != nil {
			log.Println(err)
		}
		collection.InsertOne(context.TODO(), val)
		msg.Ack()
	})

	if err != nil {
		return nil, err
	}

	return cc, nil
}
