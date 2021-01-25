package events

import (
	"encoding/json"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// ChatTeamEvent is created when somone starts a kickvote
type ChatTeamEvent struct {
	BaseEvent
	Source dto.Player `json:"source_player,omitempty"`
	Team   int        `json:"team,omitempty"`
	Text   string     `json:"text,omitempty"`
}

// Marshal creates a json string from the current struct
func (cte *ChatTeamEvent) Marshal() string {
	b, _ := json.Marshal(cte)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (cte *ChatTeamEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), cte)
}

// NewChatTeamEvent create an empty event with a proper event type
func NewChatTeamEvent() ChatTeamEvent {
	event := ChatTeamEvent{}
	event.Type = TypeChatTeam
	return event
}
