package main

import (
	"log"

	"github.com/joaovds/learn-events/pkg/events"
)

func main() {
	event := events.NewEvent("hello", "Hello, World!")

	eventHandler := events.NewEventHandler(1)

	eventDispatcher := events.NewEventDispatcher()
	err := eventDispatcher.Register(event.GetName(), eventHandler)
	if err != nil {
		log.Fatal(err)
	}

	err = eventDispatcher.Dispatch(event)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Event dispatched")
}
