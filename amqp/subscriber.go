package amqp

import (
	"context"
	"fmt"

	"github.com/streadway/amqp"
)

// Subscriber wraps the amqp client that is subscribed to one or more specific queues
// You can subscribe to different topics, but MUST NOT subscribe to the same topic multiple times,
// e.g. once with the Next method and another time with the NextromMany method.
// If you want overlapping queue subscriptions, create a second subscriber
type Subscriber struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	ctx     context.Context
	cancel  func()
}

// Close closes the channel, the connection and the go channels that pass the data.
func (s *Subscriber) Close() error {
	defer s.cancel()
	err := s.channel.Close()
	if err != nil {
		return err
	}
	return s.conn.Close()
}

// CreateExchange creates a new durable exchange
func (s *Subscriber) CreateExchange(exchange string) error {
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
	_, err := s.channel.QueueDeclare(
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

// NewSubscriber creates and starts a new subscriber that can Consume from a single of ConsumeMany queues.
// address has the format: localhost:5672
// You can subscribe to different topics, but MUST NOT subscribe to the same topic multiple times,
// e.g. once with the Next method and another time with the NextromMany method.
// If you want overlapping queue subscriptions, create a second subscriber
func NewSubscriber(address, username, password string, vhost ...string) (*Subscriber, error) {
	vhoststr := ""
	if len(vhost) > 0 {
		vhoststr = vhost[0]
	}
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, address, vhoststr))
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	subsciber := &Subscriber{
		conn:    conn,
		channel: ch,
		ctx:     ctx,
		cancel:  cancel,
	}
	return subsciber, nil
}
