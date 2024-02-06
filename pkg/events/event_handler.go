package events

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/joaovds/learn-events/pkg/rabbitmq"
)

type EventHandler struct {
	ID int
}

func NewEventHandler(id int) *EventHandler {
	return &EventHandler{
		ID: id,
	}
}

func (h *EventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		log.Panic(err)
	}
	defer ch.Close()

	ctx := context.Background()

	jsonData, err := json.Marshal(event.GetPayload())
	if err != nil {
		log.Fatal(err)
	}

	err = rabbitmq.Publish(ctx, ch, []byte(string(jsonData)), "amq.direct")
	if err != nil {
		log.Panic(err)
	}
}
