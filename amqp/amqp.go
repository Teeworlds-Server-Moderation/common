package amqp

import (
	"time"
)

var (
	// ReconnectTimeout is the default value used to reconnect until the subscribers and publishers return an error
	ReconnectTimeout = time.Minute

	// InitialReconnectTimeout allows the constructor methods to also take some time before they
	// abort their reconnection attempts
	InitialReconnectTimeout = time.Minute * 2

	// ReconnectDelay is the time waited before attempting another reconnect
	ReconnectDelay = time.Second
)
