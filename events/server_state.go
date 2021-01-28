package events

import (
	"encoding/json"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// ServerStateEvent is created when either someone requests a server state
// or when a server's server state changes.
type ServerStateEvent struct {
	BaseEvent
	dto.ServerState
}

// Marshal creates a json string from the current struct
func (sse *ServerStateEvent) Marshal() string {
	b, _ := json.Marshal(sse)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (sse *ServerStateEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), sse)
}

// NewServerStateEvent create an empty event with a proper event type
func NewServerStateEvent() ServerStateEvent {
	event := ServerStateEvent{}
	event.Type = TypeServerState
	return event
}
