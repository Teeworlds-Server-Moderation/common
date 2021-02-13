package amqp

import "github.com/streadway/amqp"

// common queue creation that is used for publisher and subscriber
// if this function is different for any of the two, you may get problems with the protocol stack of amqp,
// as you will get an error when declaring a que differently if a queue already exists.
// Idempotence is only guaranteed if created and to be created queue configurations match.
func createQueuesIfNotExists(channel *amqp.Channel, queues ...string) error {

	for _, queue := range queues {
		// declare queue if unknown
		_, err := channel.QueueDeclare(
			queue, // name
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		return err
	}
	return nil
}
