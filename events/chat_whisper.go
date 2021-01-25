package events

import (
	"encoding/json"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// ChatWhisperEvent is created when somone starts a kickvote
type ChatWhisperEvent struct {
	BaseEvent
	Source dto.Player `json:"source_player,omitempty"`
	Target dto.Player `json:"target_player,omitempty"`
	Text   string     `json:"text,omitempty"`
}

// Marshal creates a json string from the current struct
func (cwe *ChatWhisperEvent) Marshal() string {
	b, _ := json.Marshal(cwe)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (cwe *ChatWhisperEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), cwe)
}

// NewChatWhisperEvent create an empty event with a proper event type
func NewChatWhisperEvent() ChatWhisperEvent {
	event := ChatWhisperEvent{}
	event.Type = TypeChatWhisper
	return event
}
