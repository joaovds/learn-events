package main

import (
	"html/template"
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

	mux.HandleFunc("/", renderHTML)
	mux.HandleFunc("/send-message", sendMassage)

	server.ListenAndServe()
}

func renderHTML(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./web/event.html"))
	t.Execute(w, events.Event{})
}

func sendMassage(w http.ResponseWriter, r *http.Request) {
	var event events.Event
	event.Name = "hello"
	event.Payload = r.FormValue("payload")

	eventHandler := events.NewEventHandler(1)

	eventDispatcher := events.NewEventDispatcher()
	err := eventDispatcher.Register(event.GetName(), eventHandler)
	if err != nil {
		log.Fatal(err)
	}

	err = eventDispatcher.Dispatch(&event)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Event dispatched")
}
