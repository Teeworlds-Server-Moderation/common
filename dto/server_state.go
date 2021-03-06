package dto

import "encoding/json"

// ServerState is a representation of the current server state
// This dto can be sent every time something happens to the server state.
type ServerState struct {
	Map     string   `json:"map,omitempty"`
	Players []Player `json:"players"`
}

// Marshal creates a json string from the current struct
func (ss *ServerState) Marshal() string {
	b, _ := json.Marshal(ss)
	return string(b)
}

// Unmarshal fills the current struct with the unmarshalled values
func (ss *ServerState) Unmarshal(payload string) error {
	return json.Unmarshal([]byte(payload), ss)
}

// GetPlayerByID returns the requested ID
func (ss *ServerState) GetPlayerByID(ID int) (Player, error) {
	for _, player := range ss.Players {
		if player.ID == ID {
			return player, nil
		}
	}
	return Player{}, ErrIDNotFound

}
