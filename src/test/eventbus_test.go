package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventBus(t *testing.T) {
	assert := assert.New(t)

	eventBus := EventBus{}
	eventBus.publish("topic-1", "event_id_1")
	event := eventBus.pull("topic-1", "subscriber-1")

	assert.Equal("event_id_1", event)
}
