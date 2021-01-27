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
	Kills   int    `json:"kills,omitempty"`
	Deaths  int    `json:"deaths,omitempty"`
	Score   int    `json:"score,omitempty"`
	Wins    int    `json:"wins,omitempty"`
}

// PlayersSortByID sort a list of player by their ID
type PlayersSortByID []Player

func (a PlayersSortByID) Len() int           { return len(a) }
func (a PlayersSortByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PlayersSortByID) Less(i, j int) bool { return a[i].ID < a[j].ID }
