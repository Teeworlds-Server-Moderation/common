package events

import (
	"encoding/json"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// VoteKickStartedEvent is created when somone starts a kickvote
type VoteKickStartedEvent struct {
	BaseEvent
	Source dto.Player `json:"source,omitempty"`
	Target dto.Player `json:"target,omitempty"`
	// is set to true if the vote is only a visible try in the logs
	// but did not stat an actual kickvote
	Try    bool   `json:"try,omitempty"`
	Reason string `json:"reason,omitempty"`
	Forced bool   `json:"forced,omitempty"`
}

// Marshal creates a json string from the current struct
func (vkse *VoteKickStartedEvent) Marshal() string {
	b, _ := json.Marshal(vkse)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (vkse *VoteKickStartedEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), vkse)
}

// NewVoteKickStartedEvent create an empty event with a proper event type
func NewVoteKickStartedEvent() VoteKickStartedEvent {
	event := VoteKickStartedEvent{}
	event.BaseEvent = NewBaseEventTimestamped(TypeVoteKickStarted)
	return event
}
