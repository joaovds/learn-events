package main

import (
	"context"
	"log"

	"github.com/joaovds/learn-events/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		log.Panic(err)
	}
	defer ch.Close()

	ctx := context.Background()

	err = rabbitmq.Publish(ctx, ch, []byte("Hello, RabbitMQ!"), "amq.direct")
	if err != nil {
		log.Panic(err)
	}
}
