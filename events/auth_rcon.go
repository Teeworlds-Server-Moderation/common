package events

import (
	"encoding/json"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// AuthRconEvent is created when somone successfully authenticates on the econ.
type AuthRconEvent struct {
	BaseEvent
	dto.Player
	Level string `json:"auth_level,omitempty"`
	User  string `json:"credensials_user,omitempty"`
}

// Marshal creates a json string from the current struct
func (are *AuthRconEvent) Marshal() string {
	b, _ := json.Marshal(are)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (are *AuthRconEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), are)
}

// NewAuthRconEvent create an empty event with a proper event type
func NewAuthRconEvent() AuthRconEvent {
	event := AuthRconEvent{}
	event.Type = TypeAuthRcon
	return event
}
