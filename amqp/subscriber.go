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
	conn        *amqp.Connection
	channel     *amqp.Channel
	ctx         context.Context
	cancel      func()
	knownQueues map[string]bool
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

// Consume returns a channel that can be used to retrieve all received messages fro the passed queue
func (s *Subscriber) Consume(queue string) (<-chan amqp.Delivery, error) {
	if ok := s.knownQueues[queue]; ok {
		panic(fmt.Sprintf("Must not resubscribe to the same queue again: %s", queue))
	}
	s.knownQueues[queue] = true

	if err := s.createQueuesIfNotExists(queue); err != nil {
		return nil, err
	}
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

// ConsumeMany subscribes to queues and return results in the returned channel
// this should be used if more than one queue is being subscribed to, as the multiplexing of messages
// creates an overhead of n+1 extra goroutines
func (s *Subscriber) ConsumeMany(queues ...string) (<-chan amqp.Delivery, error) {

	// declare and cache
	if err := s.createQueuesIfNotExists(queues...); err != nil {
		return nil, err
	}

	deliveryChannels := make([]<-chan amqp.Delivery, 0, len(queues))
	for _, queue := range queues {
		if ok := s.knownQueues[queue]; ok {
			panic(fmt.Sprintf("Must not resubscribe to the same queue again: %s", queue))
		}
		s.knownQueues[queue] = true

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
		deliveryChannels = append(deliveryChannels, ch)
	}

	// out will receive all messages from all subscription queues
	out := make(chan amqp.Delivery, len(queues))

	// create a goroutine per subscription that pushes received messages
	// into the out channel, fan in
	for _, ch := range deliveryChannels {

		// for every created subscription channel do create a routine
		go func(ctx context.Context, channel <-chan amqp.Delivery, out chan<- amqp.Delivery) {
			for {
				select {
				case msg, ok := <-channel:
					if !ok {
						return
					}
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

func (s *Subscriber) createQueuesIfNotExists(queues ...string) error {
	return createQueuesIfNotExists(s.channel, queues...)
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
		conn:        conn,
		channel:     ch,
		ctx:         ctx,
		cancel:      cancel,
		knownQueues: make(map[string]bool),
	}
	return subsciber, nil
}
