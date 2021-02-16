package amqp

import (
	"fmt"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

// Subscriber wraps the amqp client that is subscribed to one or more specific queues
// You can subscribe to different topics, but MUST NOT subscribe to the same topic multiple times,
// e.g. once with the Next method and another time with the NextromMany method.
// If you want overlapping queue subscriptions, create a second subscriber
type Subscriber struct {
	subscriberBase
	mu sync.Mutex
}

// when reconnecting only this part is copied, as the mutex must not be touched
type subscriberBase struct {
	address  string
	username string
	password string
	vhost    string
	conn     *amqp.Connection
	channel  *amqp.Channel
}

// Close closes the channel, the connection and the go channels that pass the data.
func (s *Subscriber) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.unguardedClose()
}

func (s *Subscriber) unguardedClose() error {

	err := s.channel.Close()
	if err != nil {
		return err
	}
	return s.conn.Close()
}

// CreateExchange creates a new durable exchange
func (s *Subscriber) CreateExchange(exchange string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.channel.ExchangeDeclare(
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

	err = s.reconnect(ReconnectTimeout)
	if err != nil {
		return err
	}
	// reconnect succeeded
	return s.channel.ExchangeDeclare(
		exchange,
		"fanout",
		true,
		false,
		false,
		true,
		nil,
	)
}

// CreateQueue creates a queue that is initially bound to the default exchange
func (s *Subscriber) CreateQueue(queue string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.channel.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err == nil {
		return nil
	}

	err = s.reconnect(ReconnectTimeout)
	if err != nil {
		return err
	}
	// reconnect succeeded

	_, err = s.channel.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

// BindQueue to an exchange
func (s *Subscriber) BindQueue(queue, exchange string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.channel.QueueBind(
		queue,
		"",
		exchange,
		false,
		nil,
	)
	if err == nil {
		return nil
	}

	err = s.reconnect(ReconnectTimeout)
	if err != nil {
		return err
	}
	// reconnect succeeded
	return s.channel.QueueBind(
		queue,
		"",
		exchange,
		false,
		nil,
	)
}

// Consume returns a channel that can be used to retrieve all received messages fro the passed queue
func (s *Subscriber) Consume(queue string) (<-chan amqp.Delivery, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, err := s.channel.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err == nil {
		return c, nil
	}

	err = s.reconnect(ReconnectTimeout)
	if err != nil {
		return nil, err
	}
	// reconnect succeeded
	return s.channel.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}

func (s *Subscriber) reconnect(timeout time.Duration) error {
	end := time.Now().Add(timeout)

	var err error
	var newSub *Subscriber
	for time.Now().Before(end) {
		newSub, err = NewSubscriber(s.address, s.username, s.password, s.vhost)
		if err != nil {
			time.Sleep(ReconnectDelay)
			continue
		}
		s.unguardedClose() // ignore errors
		s.subscriberBase = *&newSub.subscriberBase
		return nil
	}
	// failed
	return err
}

// NewSubscriber creates and starts a new subscriber that can Consume from a single of ConsumeMany queues.
// address has the format: localhost:5672
// You can subscribe to different topics, but MUST NOT subscribe to the same topic multiple times,
// e.g. once with the Next method and another time with the NextromMany method.
// If you want overlapping queue subscriptions, create a second subscriber
func NewSubscriber(address, username, password string, vhosts ...string) (*Subscriber, error) {
	vhoststr := ""
	if len(vhosts) > 0 {
		vhoststr = vhosts[0]
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

	subsciber := &Subscriber{
		subscriberBase: subscriberBase{
			address:  address,
			username: username,
			password: password,
			vhost:    vhoststr,
			conn:     conn,
			channel:  ch,
		},
	}
	return subsciber, nil
}
