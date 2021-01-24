package events

import "encoding/json"

// ExecCommand is used to request a command execution
type ExecCommandEvent struct {
	BaseEvent
	User    string
	Command string
}

// Marshal creates the proper JSON string representation of the current event
func (ece *ExecCommandEvent) Marshal() string {
	b, _ := json.Marshal(ece)
	return string(b)
}
