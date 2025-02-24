package events

import (
	"authors/model"
	"authors/service"
	"log"
)

func ConsumeEvent(payload model.Payload) {
	data, event := payload.Data, payload.Event
	switch event {
	case "UPDATE_AUTHOR":
		service.UpdateAuthorWithArticle(data)
	default:
		log.Printf("Unknown event: %s", event)
	}
}
