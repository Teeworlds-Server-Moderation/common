package events

import (
	"encoding/json"
)

// MapChangedEvent is created when the map is actually changed after an option vote passes
type MapChangedEvent struct {
	BaseEvent
	OldMap string `json:"old_map,omitempty"`
	NewMap string `json:"new_map,omitempty"`
}

// Marshal creates a json string from the current struct
func (mce *MapChangedEvent) Marshal() string {
	b, _ := json.Marshal(mce)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (mce *MapChangedEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), mce)
}

// NewMapChangedEvent create an empty event with a proper event type
func NewMapChangedEvent() MapChangedEvent {
	event := MapChangedEvent{}
	event.BaseEvent = NewBaseEventTimestamped(TypeMapChanged)
	return event
}
