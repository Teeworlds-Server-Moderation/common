package events

import (
	"encoding/json"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// ChatEvent is created when somone starts a kickvote
type ChatEvent struct {
	BaseEvent
	Source dto.Player `json:"source_player,omitempty"`
	Text   string     `json:"text,omitempty"`
}

// Marshal creates a json string from the current struct
func (ce *ChatEvent) Marshal() string {
	b, _ := json.Marshal(ce)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (ce *ChatEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), ce)
}

// NewChatEvent create an empty event with a proper event type
func NewChatEvent() ChatEvent {
	event := ChatEvent{}
	event.Type = TypeChat
	return event
}
