package events

import (
	"encoding/json"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// PlayerKickedEvent is created when somone starts a kickvote
type PlayerKickedEvent struct {
	BaseEvent
	dto.Player
	Reason string `json:"reason,omitempty"`
}

// Marshal creates a json string from the current struct
func (pke *PlayerKickedEvent) Marshal() string {
	b, _ := json.Marshal(pke)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (pke *PlayerKickedEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), pke)
}

// NewPlayerKickedEvent create an empty event with a proper event type
func NewPlayerKickedEvent() PlayerKickedEvent {
	event := PlayerKickedEvent{}
	event.Type = TypePlayerKicked
	return event
}
