package amqp

import (
	"fmt"
	"strings"

	"github.com/houseofcat/turbocookedrabbit/v2/pkg/tcr"
)

// Subscriber wraps the amqp client that is subscribed to one or more specific queues
// You can subscribe to different topics, but MUST NOT subscribe to the same topic multiple times,
// e.g. once with the Next method and another time with the NextromMany method.
// If you want overlapping queue subscriptions, create a second subscriber
type Subscriber struct {
	cp        *tcr.ConnectionPool
	top       *tcr.Topologer
	consumers []*tcr.Consumer
}

// Close closes the channel, the connection and the go channels that pass the data.
func (s *Subscriber) Close() error {
	defer s.cp.Shutdown()
	var finalErr error
	for _, c := range s.consumers {
		err := c.StopConsuming(false, true)
		if err != nil {
			finalErr = err
		}
	}

	return finalErr
}

// CreateExchange creates a new durable exchange
func (s *Subscriber) CreateExchange(exchange string) error {
	return s.top.CreateExchange(
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

// CreateQueue creates a queue that is initially bound to the default exchange
func (s *Subscriber) CreateQueue(queue string) error {
	//QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args Table) (Queue, error)
	return s.top.CreateQueue(
		queue,
		false,
		true,
		false,
		false,
		false,
		nil,
	)
}

// DeleteQueue creates a queue that is initially bound to the default exchange
func (s *Subscriber) DeleteQueue(queue string) error {
	_, err := s.top.QueueDelete(
		queue,
		true,
		false,
		false,
	)
	return err
}

// BindQueue to an exchange
func (s *Subscriber) BindQueue(queue, exchange string) error {
	return s.top.QueueBind(
		&tcr.QueueBinding{
			QueueName:    queue,
			ExchangeName: exchange,
			RoutingKey:   "",
			NoWait:       false,
			Args:         nil,
		},
	)
}

// Consume returns a channel that can be used to retrieve all received messages fro the passed queue
func (s *Subscriber) Consume(queue string) (<-chan *tcr.ReceivedMessage, error) {

	consumer, err := tcr.NewConsumer(&tcr.RabbitSeasoning{},
		s.cp,
		queue,
		"",
		true,
		false,
		false,
		nil,
		0,
		0,
		0,
	)
	if err != nil {
		return nil, err
	}
	consumer.StartConsuming()

	s.consumers = append(s.consumers, consumer)
	return consumer.ReceivedMessages(), nil
}

// NewSubscriber creates and starts a new subscriber that can Consume from a single of ConsumeMany queues.
// address has the format: localhost:5672
// You can subscribe to different topics, but MUST NOT subscribe to the same topic multiple times,
// e.g. once with the Next method and another time with the NextromMany method.
// If you want overlapping queue subscriptions, create a second subscriber
func NewSubscriber(address, username, password string, vhost ...string) (*Subscriber, error) {
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

	subsciber := &Subscriber{
		cp,
		top,
		make([]*tcr.Consumer, 1),
	}
	return subsciber, nil
}
