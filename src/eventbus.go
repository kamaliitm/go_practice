// Write the publish and pull methods of an event bus with different topics and subscribers
package main

const EVENTBUS_NULL = "null"

type EventBus struct {
	partitions  map[string][]string
	subscribers map[string]map[string]int // map[subscriberId]map[topic]index
}

// add an event to the bus for a topic
func (eventbus *EventBus) publish(topic string, eventId string) {
	if eventbus.partitions == nil {
		eventbus.partitions = map[string][]string{}
	}

	if _, ok := eventbus.partitions[topic]; ok {
		eventbus.partitions[topic] = append(eventbus.partitions[topic], eventId)
	} else {
		eventbus.partitions[topic] = []string{eventId}
	}
}

// Get the next event for this subscriber from the bus for a topic
// Get the first event if this subscriber hasn't registered yet
// If the topic is empty or the subscriber has pulled all events on this topic up to this point, return NULL
func (eventbus *EventBus) pull(topic string, subscriberId string) string {
	if eventbus.subscribers == nil {
		eventbus.subscribers = map[string]map[string]int{}
	}

	if eventbus.partitions[topic] == nil {
		return EVENTBUS_NULL
	}

	if subscribedTopics, ok := eventbus.subscribers[subscriberId]; !ok {
		eventbus.subscribers[subscriberId] = map[string]int{}
		eventbus.subscribers[subscriberId][topic] = 0
		return eventbus.partitions[topic][0]
	} else {
		if index, ok2 := subscribedTopics[topic]; ok2 {
			if index < len(eventbus.partitions[topic])-1 {
				eventbus.subscribers[subscriberId][topic] = index + 1
				return eventbus.partitions[topic][index+1]
			}
		} else {
			eventbus.subscribers[subscriberId][topic] = 0
			return eventbus.partitions[topic][0]
		}
	}

	return EVENTBUS_NULL
}
