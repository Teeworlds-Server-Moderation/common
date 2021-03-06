package events

import (
	"encoding/json"
)

// MapChanedEvent is created when the map is actually changed after an option vote passes
type MapChanedEvent struct {
	BaseEvent
	OldMap string `json:"old_map,omitempty"`
	NewMap string `json:"new_map,omitempty"`
}

// Marshal creates a json string from the current struct
func (mce *MapChanedEvent) Marshal() string {
	b, _ := json.Marshal(mce)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (mce *MapChanedEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), mce)
}

// NewMapChanedEvent create an empty event with a proper event type
func NewMapChanedEvent() MapChanedEvent {
	event := MapChanedEvent{}
	event.BaseEvent = NewBaseEventTimestamped(TypeMapChanged)
	return event
}
