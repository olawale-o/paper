package events

import (
	"articles/model"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func PublishCommentEvent(payload model.RequestPayload) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	http.Post("http://localhost:7000/api/v1/comments/app-event", "application/json", bytes.NewBuffer(payloadJson))
}

func PublishAuthorEvent(payload model.RequestPayload) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	http.Post("http://localhost:7000/api/v1/authors/app-event", "application/json", bytes.NewBuffer(payloadJson))
}
