package events

import "encoding/json"

// BaseEvent contains the common fields that all events have.
type BaseEvent struct {
	// Type of the event
	Type string `json:"type,omitempty"`

	// Source server address:port that has created the event
	// Every server has its own topic that it is listening on, where
	// server
	// This is usually "<econ ip>:<econ port>"
	EventSource string `json:"event_source,omitempty"`

	// When was the event created
	Timestamp string `json:"timestamp,omitempty"`
}

// Unmarshal can be used on any event that is a composition of the BaseEvent
// in order to distinguish between different event types that are
// correctly retrieved by this function.
// Depending on the type, one can properly unmarshal the corresponding event type
func (be *BaseEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), be)
}

// SetEventSource allows to set the event's source, e.g. which server/econ connection created this event
func (be *BaseEvent) SetEventSource(eventSource string) {
	be.EventSource = eventSource
}

// NewBaseEventTimestamped creates a new BaseEvent that is timestamped with time.Now()
// and has the passed eventType as type parameter. The EventSource is left empty.
func NewBaseEventTimestamped(eventType string) BaseEvent {
	return BaseEvent{
		Type:      eventType,
		Timestamp: FormatedTimestamp(),
	}
}
