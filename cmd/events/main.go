package main

import (
	"log"
	"net/http"

	"github.com/joaovds/learn-events/pkg/events"
)

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()

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
