package events

import (
	"encoding/json"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// RequestServerStateEvent is created when either someone requests a server state
// or when a server's server state changes.
type RequestServerStateEvent struct {
	BaseEvent
	dto.ServerState
}

// Marshal creates a json string from the current struct
func (rss *RequestServerStateEvent) Marshal() string {
	b, _ := json.Marshal(rss)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (rss *RequestServerStateEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), rss)
}

// NewRequestServerStateEvent create an empty event with a proper event type
func NewRequestServerStateEvent() RequestServerStateEvent {
	event := RequestServerStateEvent{}
	event.Type = TypeRequestServerState
	return event
}
