package rabbitmq

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}

	return ch, nil
}

func Consume(ch *amqp.Channel, msgsOut chan<- amqp.Delivery) error {
	msgs, err := ch.Consume(
		"my_queue",
		"my_consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		msgsOut <- msg
	}

	return nil
}

func Publish(ctx context.Context, ch *amqp.Channel, msg []byte, exName string) error {
	err := ch.PublishWithContext(
		ctx,
		exName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
