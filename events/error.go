package events

import (
	"encoding/json"
)

// ErrorEvent is returned when some errors occur in between specific states
type ErrorEvent struct {
	BaseEvent
	Error string `json:"error, omitempty"`
}

// Marshal creates a json string from the current struct
func (ee *ErrorEvent) Marshal() string {
	b, _ := json.Marshal(ee)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (ee *ErrorEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), ee)
}

// NewErrorEvent creates a new ErrorEvent that is requrned when for example
// requests cannot be answered, as a player is not online anymore in order to request their
func NewErrorEvent() ErrorEvent {
	event := ErrorEvent{}
	event.Type = TypeError
	return event
}
