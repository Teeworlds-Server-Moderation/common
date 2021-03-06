package events

import (
	"encoding/json"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// PlayerLeftEvent is created when a player leaves some server.
type PlayerLeftEvent struct {
	BaseEvent
	dto.Player
	Reason string `json:"reason,omitempty"`
}

// Marshal creates a json string from the current struct
func (ple *PlayerLeftEvent) Marshal() string {
	b, _ := json.Marshal(ple)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (ple *PlayerLeftEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), ple)
}

// NewPlayerLeftEvent create an empty event with a proper event type
func NewPlayerLeftEvent() PlayerLeftEvent {
	event := PlayerLeftEvent{}
	event.BaseEvent = NewBaseEventTimestamped(TypePlayerLeft)
	return event
}
