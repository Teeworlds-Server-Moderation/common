package dto

import "errors"

// dto - Data Transfer Objects
// These objects are used to transfer data between
// A requestor and a requestor.
// The semantics for tranferring these structs is the same as with evets,
// but contrary to events these are not created, but requested.
// For example one may request the current server state or the current banlist,voteban,mute list

var (
	// ErrIDNotFound is returned by functions that are passed an ID, but that ID cannot be found
	// be it the player in the server state of some IP in an ID to mapping.
	ErrIDNotFound = errors.New("error, could not find the requested id")
)

const (
	// SanctionTypeBan is the type ban, where a player cannot rejoin a server for a specific time.
	SanctionTypeBan = "ban"

	// SanctionTypeMute is the mute of a player that is not allowed to chat for a specific period of time.
	SanctionTypeMute = "mute"

	// SanctionTypeVoteBan is the sanction of a player not being allowed to start either spec- or kick-votes.
	SanctionTypeVoteBan = "vote_ban"

	// SanctionTypeTrollPit is the sanction where a player is moved to a different chat space where they cannot interact
	// with normal players anymore via the chat, Imagine a pit of trolls only being able to talk to other trolls in their same pit
	// while other players can talk to them from above, but cannot hear what the trolls are saying.
	SanctionTypeTrollPit = "troll_pit"
)
