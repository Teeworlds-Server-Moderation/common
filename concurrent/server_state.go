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
	Map     string
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

// PlayerLeaveAll removes all players that are currently on the server from the server
// and returns a list with these players.
// This function is useful for when a map change happens and all players are disconnected for
// a short amount of time and during this time they may leave the server, as they do not like the map.
// In order to catch that state, players should be removed from the online player list during map changes
func (ss *ServerState) PlayerLeaveAll() []dto.Player {
	ss.mu.Lock()
	defer ss.mu.Unlock()

	players := make([]dto.Player, 0, len(ss.players))
	for id, player := range ss.players {
		players = append(players, player)
		delete(ss.players, id)
	}

	sort.Sort(dto.PlayersSortByID(players))
	return players
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
		Map:     ss.Map,
		Players: players,
	}
}

// GetPlayer fetches a specific Player state from the current server state.
func (ss *ServerState) GetPlayer(ID int) dto.Player {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	return ss.players[ID]
}

// GetMap returns the currently played map. If it's empty, the map has not yet been changed in
// order to be visible in here.
func (ss *ServerState) GetMap() string {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	return ss.Map
}

// SetMap updats the currently played map on the server.
func (ss *ServerState) SetMap(newMap string) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	ss.Map = newMap
}
