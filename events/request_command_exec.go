package events

import "encoding/json"

// RequestCommandExecEvent is used to request a command execution
type RequestCommandExecEvent struct {
	BaseEvent
	Requestor string `json:"requestor,omitempty"`
	Command   string `json:"command,omitempty"`
}

// Marshal creates the proper JSON string representation of the current event
func (rce *RequestCommandExecEvent) Marshal() string {
	b, _ := json.Marshal(rce)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (rce *RequestCommandExecEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), rce)
}

// NewRequestCommandExecEvent create an empty event with a proper event type
func NewRequestCommandExecEvent() RequestCommandExecEvent {
	event := RequestCommandExecEvent{}
	event.Type = TypeRequestCommandExec
	return event
}
