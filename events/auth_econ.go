package events

import "encoding/json"

// AuthEconEvent is created when somone successfully authenticates on the econ.
type AuthEconEvent struct {
	BaseEvent
	IP   string
	Port int
	ID   int
}

// Marshal creates a json string from the current struct
func (aee *AuthEconEvent) Marshal() string {
	b, _ := json.Marshal(aee)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (aee *AuthEconEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), aee)
}

// NewAuthEconEvent create an empty event with a proper event type
func NewAuthEconEvent() AuthEconEvent {
	event := AuthEconEvent{}
	event.Type = TypeAuthEcon
	return event
}
