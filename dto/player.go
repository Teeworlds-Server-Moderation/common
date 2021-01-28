package dto

import "encoding/json"

// Player is the object that represents a player
//  with all of their information.
type Player struct {
	Name    string `json:"name,omitempty"`
	Clan    string `json:"clan,omitempty"`
	IP      string `json:"ip,omitempty"`
	Port    int    `json:"port,omitempty"`
	ID      int    `json:"id,omitempty"`
	Country int    `json:"country,omitempty"`
	Version int    `json:"version,omitempty"`
	Kills   int    `json:"kills,omitempty"`
	Deaths  int    `json:"deaths,omitempty"`
	Score   int    `json:"score,omitempty"`
	Wins    int    `json:"wins,omitempty"`
}

// Marshal creates a json string from the current struct
func (p *Player) Marshal() string {
	b, _ := json.Marshal(p)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (p *Player) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), p)
}
