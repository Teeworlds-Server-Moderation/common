package events

import "encoding/json"

// PlayerJoinEvent is created when a player joins some server.
type PlayerJoinEvent struct {
	BaseEvent
	Name    string `json:"name,omitempty"`
	Clan    string `json:"clan,omitempty"`
	IP      string `json:"ip,omitempty"`
	Port    int    `json:"port,omitempty"`
	ID      int    `json:"id,omitempty"`
	Country int    `json:"country,omitempty"`
	Version int    `json:"version,omitempty"`
}

// NewEventPlayerJoin creates a new PlayerJoinEvent
func NewEventPlayerJoin(source, timestamp, name, clan, ip string, port, id, country, version int) PlayerJoinEvent {
	return PlayerJoinEvent{
		BaseEvent: BaseEvent{
			Type:      TypePlayerJoin,
			Source:    source,
			Timestamp: timestamp,
		},
		Name:    name,
		Clan:    clan,
		IP:      ip,
		Port:    port,
		ID:      id,
		Country: country,
		Version: version,
	}
}

// Marshal creates a json string from the current struct
func (pje *PlayerJoinEvent) Marshal() string {
	b, _ := json.Marshal(pje)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (pje *PlayerJoinEvent) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), pje)
}
