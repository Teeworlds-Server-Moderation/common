package dto

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
}
