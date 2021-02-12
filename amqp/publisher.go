package amqp

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Publisher struct {
	queue          string
	conn           *amqp.Connection
	channel        *amqp.Channel
	declaredQueues map[string]bool
}

// Close waits a second and then closes the client connection as well as the subsciber
// and all internally used channels
func (p *Publisher) Close() error {
	if err := p.channel.Close(); err != nil {
		return err
	}
	return p.conn.Close()
}

// PublishTo allows to specify a different topic other than the default one.
func (p *Publisher) PublishTo(queue string, msg interface{}) error {
	if err := p.createQueuesIfNotExists(queue); err != nil {
		return err
	}
	return p.channel.Publish(
		"",    // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(toString(msg)),
		})
}

func (p *Publisher) createQueuesIfNotExists(queues ...string) error {
	return createQueuesIfNotExists(p.declaredQueues, p.channel, queues...)
}

// NewPublisher creates and starts a new Publisher that receives new messages via
// a string channel that can be
// address has the format: localhost:5672
func NewPublisher(address, username, password string, vhost ...string) (*Publisher, error) {
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

	publisher := &Publisher{
		declaredQueues: make(map[string]bool, 1), // you want to at least publish to one queue
		conn:           conn,
		channel:        ch,
	}

	return publisher, err
}
