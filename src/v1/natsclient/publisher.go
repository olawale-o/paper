package natsclient

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func PublishMessage(ctx context.Context, subject string, data []byte) {
	nc, _ := nats.Connect(nats.DefaultURL)
	err := nc.Publish(subject, data)
	if err != nil {
		fmt.Println("error publishing message:", err)
		panic(err)
	}
	time.Sleep(2 * time.Second)
}
