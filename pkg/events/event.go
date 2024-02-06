package events

import "time"

type Event struct {
	Name    string      `json:"name"`
	Payload interface{} `json:"payload"`
}

func NewEvent(name string, payload interface{}) *Event {
	return &Event{
		Name:    name,
		Payload: payload,
	}
}

func (e *Event) GetName() string {
	return e.Name
}

func (e *Event) GetPayload() interface{} {
	return e.Payload
}

func (e *Event) GetDateTime() time.Time {
	return time.Now()
}
