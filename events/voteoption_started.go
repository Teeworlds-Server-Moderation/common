package events

import (
	"encoding/json"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// VoteOptionStartedEvent is created when somone starts a kickvote
type VoteOptionStartedEvent struct {
	BaseEvent
	Source dto.Player `json:"source,omitempty"`
	// is set to true if the vote is only a visible try in the logs
	// but did not stat an actual kickvote
	Try    bool   `json:"try,omitempty"`
	Option string `json:"option,omitempty"`
	Reason string `json:"reason,omitempty"`
	Forced bool   `json:"forced,omitempty"`
}

// Marshal creates a json string from the current struct
func (vose *VoteOptionStartedEvent) Marshal() string {
	b, _ := json.Marshal(vose)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (vose *VoteOptionStartedEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), vose)
}

// NewVoteOptionStartedEvent create an empty event with a proper event type
func NewVoteOptionStartedEvent() VoteOptionStartedEvent {
	event := VoteOptionStartedEvent{}
	event.Type = TypeVoteOptionStarted
	return event
}
