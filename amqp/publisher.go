package amqp

import (
	"fmt"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type Publisher struct {
	publisherBase
	mu sync.Mutex
}

type publisherBase struct {
	address  string
	username string
	password string
	vhost    string
	conn     *amqp.Connection
	channel  *amqp.Channel
}

// Close closes the client connection as well as the subsciber
// and all internally used channels
// Close closes the channel, the connection and the go channels that pass the data.
func (p *Publisher) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.unguardedClose()
}

func (p *Publisher) unguardedClose() error {

	err := p.channel.Close()
	if err != nil {
		return err
	}
	return p.conn.Close()
}

// CreateExchange creates a new durable exchange
func (p *Publisher) CreateExchange(exchange string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	err := p.channel.ExchangeDeclare(
		exchange,
		"fanout",
		true,
		false,
		false,
		true,
		nil,
	)
	if err == nil {
		return nil
	}

	err = p.reconnect(ReconnectTimeout)
	if err != nil {
		return err
	}
	// reconnect succeeded
	return p.channel.ExchangeDeclare(
		exchange,
		"fanout",
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
	p.mu.Lock()
	defer p.mu.Unlock()

	err := p.channel.Publish(
		exchange, // exchange
		queue,    // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(toString(msg)),
		})
	if err == nil {
		return nil
	}

	err = p.reconnect(ReconnectTimeout)
	if err != nil {
		return err
	}
	// reconnect succeeded

	return p.channel.Publish(
		exchange, // exchange
		queue,    // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(toString(msg)),
		})
}

func (p *Publisher) reconnect(timeout time.Duration) error {
	end := time.Now().Add(timeout)

	var err error
	var newPub *Publisher
	for time.Now().Before(end) {
		newPub, err = NewPublisher(p.address, p.username, p.password, p.vhost)
		if err != nil {
			time.Sleep(ReconnectDelay)
			continue
		}
		p.unguardedClose() // ignore errors
		p.publisherBase = *&newPub.publisherBase
		return nil
	}
	// failed
	return err
}

// NewPublisher creates and starts a new Publisher that receives new messages via
// a string channel that can be
// address has the format: localhost:5672
func NewPublisher(address, username, password string, vhost ...string) (*Publisher, error) {
	vhoststr := ""
	if len(vhost) > 0 {
		vhoststr = vhost[0]
	}
	var err error
	var conn *amqp.Connection
	var ch *amqp.Channel

	end := time.Now().Add(InitialReconnectTimeout)
	for time.Now().Before(end) {
		conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, address, vhoststr))
		if err != nil {
			time.Sleep(ReconnectDelay)
			continue
		}
		ch, err = conn.Channel()
		if err != nil {
			time.Sleep(ReconnectDelay)
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}

	publisher := &Publisher{
		publisherBase: publisherBase{
			address:  address,
			username: username,
			password: password,
			vhost:    vhoststr,
			conn:     conn,
			channel:  ch,
		},
	}

	return publisher, err
}
