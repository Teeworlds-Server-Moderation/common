package amqp

import (
	"fmt"
	"strings"
	"time"

	"github.com/houseofcat/turbocookedrabbit/v2/pkg/tcr"
)

const (
	deliveryModePersistent = 2
)

type Publisher struct {
	cp  *tcr.ConnectionPool
	pub *tcr.Publisher
	top *tcr.Topologer
}

// Close closes the client connection as well as the subsciber
// and all internally used channels
// Close closes the channel, the connection and the go channels that pass the data.
func (p *Publisher) Close() error {
	p.pub.Shutdown(false)
	p.cp.Shutdown()
	return nil
}

// CreateExchange creates a new durable exchange
func (p *Publisher) CreateExchange(exchange string) error {
	return p.top.CreateExchange(
		exchange,
		"fanout",
		false,
		true,
		false,
		false,
		true,
		nil,
	)
}

// Publish allows to specify a different topic other than the default one.
// leave queue empty to only send to the queue
func (p *Publisher) Publish(exchange, queue string, msg interface{}) error {
	p.pub.PublishWithConfirmation(
		&tcr.Letter{
			Body: []byte(toString(msg)),
			Envelope: &tcr.Envelope{
				Exchange:     exchange,
				RoutingKey:   queue,
				Mandatory:    false,
				Immediate:    false,
				ContentType:  "application/json",
				DeliveryMode: deliveryModePersistent,
			},
		}, time.Minute)
	receipt := <-p.pub.PublishReceipts()
	return receipt.Error
}

// NewPublisher creates and starts a new Publisher that receives new messages via
// a string channel that can be
// address has the format: localhost:5672
func NewPublisher(address, username, password string, vhost ...string) (*Publisher, error) {
	vhoststr := ""
	if len(vhost) > 0 {
		vhoststr = strings.TrimLeft(vhost[0], "/")
	}

	cp, err := tcr.NewConnectionPool(&tcr.PoolConfig{
		URI:                fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, address, vhoststr),
		MaxConnectionCount: 5,
		ConnectionTimeout:  10,
	})
	if err != nil {
		return nil, err
	}
	top := tcr.NewTopologer(cp)

	pub := tcr.NewPublisher(
		cp,
		0,
		0,
		5000*time.Second,
	)

	return &Publisher{
		cp,
		pub,
		top,
	}, nil
}
