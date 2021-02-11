package amqp

import (
	"context"
	"fmt"

	"github.com/streadway/amqp"
)

// Subscriber wraps the mqtt client that is subscribed to a specific topic
// in a pretty simple to use manner.
// initially you connect to your broker and fetch reveived messages with the method
// Next(). Next() is a blocking call that waits for a channel to contain a message or
// until the Close() method has been called that cancels an internally wrapped context, which
// immediatly terminates
type Subscriber struct {
	queue   string
	conn    *amqp.Connection
	channel *amqp.Channel
	ctx     context.Context
	cancel  func()
}

// Close waits a second and then closes the client connection as well as the subsciber
// and all internally used channels
func (s *Subscriber) Close() error {
	defer s.cancel()
	err := s.channel.Close()
	if err != nil {
		return err
	}
	return s.conn.Close()
}

// Next blocks until the next message from the broker is received
// the bool indicates that the subscriber was closed
// you can use this in a for loop until ok is false, preferrably in an own goroutine
func (s *Subscriber) Next() (<-chan amqp.Delivery, error) {
	return s.channel.Consume(
		s.queue, // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
}

func (s *Subscriber) NextFromQueues(queues ...string) (<-chan amqp.Delivery, error) {
	channels := make([]<-chan amqp.Delivery, len(queues))
	for _, queue := range queues {
		ch, err := s.channel.Consume(
			queue, // queue
			"",    // consumer
			true,  // auto-ack
			false, // exclusive
			false, // no-local
			false, // no-wait
			nil,   // args
		)
		if err != nil {
			return nil, err
		}
		channels = append(channels, ch)
	}

	// out will receive all messages from all subscription queues
	out := make(chan amqp.Delivery, len(queues))

	// create a goroutine per subscription that pushes received messages
	// into the out channel, fan in
	for _, ch := range channels {

		// for every created subscription channel do create a routine
		go func(ctx context.Context, channel <-chan amqp.Delivery, out chan amqp.Delivery) {
			for {
				select {
				case msg := <-channel:
					out <- msg
				case <-ctx.Done():
					return
				}
			}
		}(s.ctx, ch, out)
	}

	// extra routine to close 'out channel' when
	// the subscriber is closed
	go func(ctx context.Context) {
		for {
			select {
			// block until context is closed
			case <-ctx.Done():
				close(out)
				return
			}
		}
	}(s.ctx)
	return out, nil
}

func (s *Subscriber) createQueueIfNotExists(queue string) error {
	_, err := s.channel.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	return err
}

// NewSubscriber creates and starts a new subscriber that receives new messages via
// a string channel that can be
// address has the format: tcp://localhost:1883
func NewSubscriber(queue, username, password, hostname string, port int, vhost ...string) (*Subscriber, error) {
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

	ctx, cancel := context.WithCancel(context.Background())
	subsciber := &Subscriber{
		queue:   queue,
		conn:    conn,
		channel: ch,
		ctx:     ctx,
		cancel:  cancel,
	}

	if queue != "" {
		err = subsciber.createQueueIfNotExists(queue)
	}
	return subsciber, err
}
