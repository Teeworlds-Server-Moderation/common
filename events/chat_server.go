package events

import (
	"encoding/json"
)

// ChatServerEvent is created when somone starts a kickvote
type ChatServerEvent struct {
	BaseEvent
	Text string `json:"text,omitempty"`
}

// Marshal creates a json string from the current struct
func (cse *ChatServerEvent) Marshal() string {
	b, _ := json.Marshal(cse)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (cse *ChatServerEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), cse)
}

// NewChatServerEvent create an empty event with a proper event type
func NewChatServerEvent() ChatServerEvent {
	event := ChatServerEvent{}
	event.Type = TypeChatServer
	return event
}
