package concurrent

import (
	"sort"
	"sync"

	"github.com/Teeworlds-Server-Moderation/common/dto"
)

// ServerState represents the server state from the perspective of a specific server monitor that
// is monitoring an dparsing the logs of that specific server.
type ServerState struct {
	mu      sync.Mutex
	players map[int]dto.Player
}

// NewServerState create a new mutex locked server state that can be modified
// concurrently.
func NewServerState() *ServerState {
	return &ServerState{
		players: make(map[int]dto.Player, 64),
	}
}

// PlayerJoin can be used when a player joins in order to change the server state to one more player
func (ss *ServerState) PlayerJoin(ID int, player dto.Player) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	ss.players[ID] = player
}

// PlayerLeave changes the player state to one less player
func (ss *ServerState) PlayerLeave(ID int) dto.Player {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	player := ss.players[ID]
	delete(ss.players, ID)
	player.ID = ID
	return player
}

// GetState returns the current server state as a list of players sorted by their ID
func (ss *ServerState) GetState() dto.ServerState {
	ss.mu.Lock()
	players := make([]dto.Player, 0, len(ss.players))
	for _, player := range ss.players {
		players = append(players, player)
	}
	ss.mu.Unlock()

	sort.Sort(dto.PlayersSortByID(players))
	return dto.ServerState{
		Players: players,
	}
}

// GetPlayer fetches a specific Player state from the current server state.
func (ss *ServerState) GetPlayer(ID int) dto.Player {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	return ss.players[ID]
}
