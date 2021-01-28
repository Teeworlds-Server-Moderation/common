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
