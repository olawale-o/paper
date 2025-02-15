package events

import (
	"articles/model"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func PublishCommentEvent(payload model.Payload) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	http.Post("http://localhost:7000/api/v1/comments/app-event", "application/json", bytes.NewBuffer(payloadJson))
}
