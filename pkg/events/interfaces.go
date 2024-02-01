package events

import "time"

type EventInterface interface {
  GetName() string
  GetPayload() interface{}
  GetDateTime() time.Time
}

type EventHandlerInterface interface {
  Handle(event EventInterface)
}

type EventDispatcherInterface interface {
  Register(eventName string, handler EventHandlerInterface) error
  Dispatch(event EventInterface) error
  Remove(eventName string, handler EventHandlerInterface) error
  Has(eventName string, handler EventHandlerInterface) bool
  Clear() error
}