package events

// this package defines all of the existing events as well as their respective structs

const (

	// FROM SERVERS

	// TypeError is returned from server monitors when they
	// cannot answer the requests sent by the other microservices.
	TypeError = "EVENT:ERROR"
	// TypePlayerJoined is an event that is created by a monitor when a player joins a server
	TypePlayerJoined = "EVENT:PLAYER_JOIN"
	// TypePlayerLeft is created when a player leaves the server
	TypePlayerLeft = "EVENT:PLAYER_LEAVE"
	// TypeVoteKickStarted is created when a kickvote is started
	TypeVoteKickStarted = "EVENT:KICKVOTE_START"
	// TypeVoteSpecStarted is created when a specvote is started,
	// trying to move a player to the spectators
	TypeVoteSpecStarted = "EVENT:SPECVOTE_START"
	// TypeVoteOptionStarted is created when a player for example start a map vote
	TypeVoteOptionStarted = "EVENT:OPTIONVOTE_START"
	// TypeChat is created when some chat message is written by someone
	TypeChat = "EVENT:CHAT_ALL"
	// TypeChatWhisper is created when someone whispers to someone else
	TypeChatWhisper = "EVENT:CHAT_WHISPER"
	// TypeChatTeam is created when someone writes in the teamchat
	TypeChatTeam = "EVENT:CHAT_TEAM"
	// TypeChatServer is created when the server sends a message or someone talks via rcon.
	TypeChatServer = "EVENT:CHAT_SERVER"
	// TypePlayerMuted is created when a player is muted by either the server or an admin/moderator
	TypePlayerMuted = "EVENT:PLAYER_MUTED"
	// TypePlayerKicked is created when someone is kicked from the server.
	TypePlayerKicked = "EVENT:PLAYER_KICKED"
	// TypePlayerBanned is created when someone is banned from the server.
	TypePlayerBanned = "EVENT:PLAYER_BANNED"
	// TypePlayerVoteBanned is created when a player receives a voteban from an admin/moderator.
	TypePlayerVoteBanned = "EVENT_PLAYER_VOTEBANNED"
	// TypeAuthEcon is created when someone logs into econ.
	TypeAuthEcon = "EVENT:AUTH_ECON"
	// TypeAuthRcon is created when a player logs into the rcon.
	TypeAuthRcon = "EVENT:AUTH_RCON"
	// TypePlayerDied is created when a player killed another player
	TypePlayerDied = "EVENT:PLAYER_DIED"
	// TypeServerState is created when a microservice requests the state of a server
	// or when a server state change occurs.
	TypeServerState = "EVENT:SERVER_STATE"

	// TO SERVERS

	// TypeRequestCommandExec is used to forward command execution requests to the Teeworlds servers
	TypeRequestCommandExec = "REQEST:COMMAND_EXEC"
	// TypeRequestServerState is created when a specific microservice needs more data for
	// an interaction with the server. For example listing the current player list.
	TypeRequestServerState = "REQUEST:SERVER_STATE"
)
