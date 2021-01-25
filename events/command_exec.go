package events

import "encoding/json"

// CommandExecEvent is used to request a command execution
type CommandExecEvent struct {
	BaseEvent
	Requestor string `json:"requestor,omitempty"`
	Command   string `json:"command,omitempty"`
}

// Marshal creates the proper JSON string representation of the current event
func (cee *CommandExecEvent) Marshal() string {
	b, _ := json.Marshal(cee)
	return string(b)
}

// NewCommandExecEvent create an empty event with a proper event type
func NewCommandExecEvent() CommandExecEvent {
	event := CommandExecEvent{}
	event.Type = TypeCommandExec
	return event
}
