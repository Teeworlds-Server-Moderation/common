package events

import (
	"encoding/json"
	"time"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// PlayerBannedEvent is created when somone starts a kickvote
type PlayerBannedEvent struct {
	BaseEvent
	dto.Player
	Duration time.Duration `json:"duration,omitempty"`
	Reason   string        `json:"reason,omitempty"`
}

// Marshal creates a json string from the current struct
func (pbe *PlayerBannedEvent) Marshal() string {
	b, _ := json.Marshal(pbe)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (pbe *PlayerBannedEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), pbe)
}

// NewPlayerBannedEvent create an empty event with a proper event type
func NewPlayerBannedEvent() PlayerBannedEvent {
	event := PlayerBannedEvent{}
	event.BaseEvent = NewBaseEventTimestamped(TypePlayerBanned)
	return event
}
