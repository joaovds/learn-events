package main

import (
	"log"

	"github.com/joaovds/learn-events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
  ch, err := rabbitmq.OpenChannel()
  if err != nil {
    log.Panic(err)
  }
  defer ch.Close()

  msgs := make(chan amqp.Delivery)

  go rabbitmq.Consume(ch, msgs)

  for msg := range msgs {
    log.Printf("Received message: %s\n", msg.Body)
    msg.Ack(false)
  }
}
