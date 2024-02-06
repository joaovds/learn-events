package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEventMethods(t *testing.T) {
	t.Run("should return new event", func(t *testing.T) {
		event := NewEvent("test", "test")

		assert.Equal(t, "test", event.Name)
		assert.Equal(t, "test", event.Payload)
	})

	t.Run("should return event name", func(t *testing.T) {
		eventName := "test"
		event := NewEvent(eventName, "test")

		assert.Equal(t, eventName, event.GetName())
	})

	t.Run("should return event payload", func(t *testing.T) {
		payload := "test"
		event := NewEvent("test", payload)

		assert.Equal(t, payload, event.GetPayload())
	})

	t.Run("should return event date time", func(t *testing.T) {
		event := NewEvent("test", "test")

		assert.NotEmpty(t, event.GetDateTime())
		assert.IsType(t, time.Now(), event.GetDateTime())
	})
}
