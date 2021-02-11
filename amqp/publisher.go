package amqp

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Publisher struct {
	queue   string
	conn    *amqp.Connection
	channel *amqp.Channel
}

// Close waits a second and then closes the client connection as well as the subsciber
// and all internally used channels
func (p *Publisher) Close() error {
	err := p.channel.Close()
	if err != nil {
		return err
	}
	return p.conn.Close()
}

// Publish pushes the message into a channel which is emptied by a concurrent goroutine
// and published to th ebroker at the specified topic.
// use string, []byte, int, int64, float32, float64 or
// Message{
//	  Topic string,
//    Payload inteface{}
// }
// allows you to control the taget topic
// if you are using custom structs, please convert them into JSON before
// passing them to this function
func (p *Publisher) Publish(msg interface{}) error {

	switch m := msg.(type) {
	case Message:
		// if it's a message, you can explicity control the topic
		return p.channel.Publish(
			"",      // exchange
			m.Queue, // routing key
			false,   // mandatory
			false,   // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(m.Payload),
			})
	}
	return p.channel.Publish(
		"",      // exchange
		p.queue, // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(toString(msg)),
		})
}

// PublishTo allows to specify a different topic other than the default one.
func (p *Publisher) PublishTo(queue string, msg interface{}) error {
	switch m := msg.(type) {
	case Message:
		return p.channel.Publish(
			"",    // exchange
			queue, // routing key
			false, // mandatory
			false, // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(m.Payload),
			})
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

func (p *Publisher) createQueueIfNotExists(queue string) error {
	_, err := p.channel.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	return err
}

// NewPublisher creates and starts a new Publisher that receives new messages via
// a string channel that can be
// address has the format: tcp://localhost:1883
func NewPublisher(queue, username, password, hostname string, port int, vhost ...string) (*Publisher, error) {
	vhoststr := ""
	if len(vhost) > 0 {
		vhoststr = vhost[0]
	}
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/%s", username, password, hostname, port, vhoststr))
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	publisher := &Publisher{
		queue:   queue,
		conn:    conn,
		channel: ch,
	}

	if queue != "" {
		err = publisher.createQueueIfNotExists(queue)
	}
	return publisher, err
}
