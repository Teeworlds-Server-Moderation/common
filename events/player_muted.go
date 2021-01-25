package events

import (
	"encoding/json"
	"time"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// PlayerMutedEvent is created when somone starts a kickvote
type PlayerMutedEvent struct {
	BaseEvent
	dto.Player
	Duration time.Duration `json:"duration,omitempty"`
	Reason   string        `json:"reason,omitempty"`
}

// Marshal creates a json string from the current struct
func (pme *PlayerMutedEvent) Marshal() string {
	b, _ := json.Marshal(pme)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (pme *PlayerMutedEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), pme)
}

// NewPlayerMutedEvent create an empty event with a proper event type
func NewPlayerMutedEvent() PlayerMutedEvent {
	event := PlayerMutedEvent{}
	event.Type = TypeCommandExec
	return event
}
