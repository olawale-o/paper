package events

import (
	"comments/model"
	"comments/service"
	"log"
)

func ConsumeEvent(payload model.Payload) {
	data, event := payload.Data, payload.Event
	switch event {
	case "NEW_COMMENT":
		value := data.(model.ArticleData)
		service.NewComment(value)
	default:
		log.Printf("Unknown event: %s", event)
	}
}
