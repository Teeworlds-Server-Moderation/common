package events

import (
	"encoding/json"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// VoteSpecStartedEvent is created when somone starts a kickvote
type VoteSpecStartedEvent struct {
	BaseEvent
	Source dto.Player `json:"source,omitempty"`
	Target dto.Player `json:"target,omitempty"`
	Try    bool       `json:"try,omitempty"`
	Reason string     `json:"reason,omitempty"`
	Forced bool       `json:"forced,omitempty"`
}

// Marshal creates a json string from the current struct
func (vsse *VoteSpecStartedEvent) Marshal() string {
	b, _ := json.Marshal(vsse)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (vsse *VoteSpecStartedEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), vsse)
}

// NewVoteSpecStartedEvent create an empty event with a proper event type
func NewVoteSpecStartedEvent() VoteSpecStartedEvent {
	event := VoteSpecStartedEvent{}
	event.BaseEvent = NewBaseEventTimestamped(TypeVoteSpecStarted)
	return event
}
