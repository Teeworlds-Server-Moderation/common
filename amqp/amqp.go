package amqp

import (
	"fmt"
	"strings"

	"github.com/houseofcat/turbocookedrabbit/v2/pkg/tcr"
)

func newConnectionPool(username, password, address string, vhost ...string) (*tcr.ConnectionPool, error) {
	vhoststr := ""
	if len(vhost) > 0 {
		vhoststr = strings.TrimLeft(vhost[0], "/")
	}

	cp, err := tcr.NewConnectionPool(&tcr.PoolConfig{
		URI:                fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, address, vhoststr),
		MaxConnectionCount: 5,
		ConnectionTimeout:  10,
		Heartbeat:          3,
	})
	if err != nil {
		return nil, err
	}
	return cp, nil
}
