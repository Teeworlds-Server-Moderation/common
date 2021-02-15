package amqp

import (
	"time"
)

var (
	// ReconnectTimeout is the default value used to reconnect until the subscribers and publishers return an error
	ReconnectTimeout = time.Minute
)
