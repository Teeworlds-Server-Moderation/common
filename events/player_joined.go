package events

import (
	"encoding/json"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// PlayerJoinedEvent is created when a player joins some server.
type PlayerJoinedEvent struct {
	BaseEvent
	dto.Player
}

// Marshal creates a json string from the current struct
func (pje *PlayerJoinedEvent) Marshal() string {
	b, _ := json.Marshal(pje)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (pje *PlayerJoinedEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), pje)
}

// NewPlayerJoinedEvent create an empty event with a proper event type
func NewPlayerJoinedEvent() PlayerJoinedEvent {
	event := PlayerJoinedEvent{}
	event.BaseEvent = NewBaseEventTimestamped(TypePlayerJoined)
	return event
}
