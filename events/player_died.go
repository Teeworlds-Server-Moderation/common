package events

import (
	"encoding/json"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// PlayerDiedEvent is created when somone starts a kickvote
// A PlayerKilledEvent is not necessary, as this contains both of them and even more,
// especially killed by the game or the world.
type PlayerDiedEvent struct {
	BaseEvent
	Victim dto.Player `json:"victim,omitempty"`
	Killer dto.Player `json:"killer,omitempty"`
	Weapon int
}

// Marshal creates a json string from the current struct
func (pbe *PlayerDiedEvent) Marshal() string {
	b, _ := json.Marshal(pbe)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (pbe *PlayerDiedEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), pbe)
}

// NewPlayerDiedEvent create an empty event with a proper event type
func NewPlayerDiedEvent() PlayerDiedEvent {
	event := PlayerDiedEvent{}
	event.Type = TypePlayerDied
	return event
}
