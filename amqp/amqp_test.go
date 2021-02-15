package amqp

import (
	"testing"
	"time"
)

func credentials() (address, user, password string) {
	return "localhost:5672", "tw-admin", "unsecure_password"
}

func TestPubSub(t *testing.T) {
	pub, err := NewPublisher(credentials())
	if err != nil {
		t.Fatal(err)
	}
	defer pub.Close()
	sub, err := NewSubscriber(credentials())
	if err != nil {
		t.Fatal(err)
	}
	defer sub.Close()

	exchange := "exchange"
	err = pub.CreateExchange(exchange)
	if err != nil {
		t.Fatal(err)
	}

	queue := "consumerQueue"
	err = sub.CreateQueue(queue)
	if err != nil {
		t.Fatal(err)
	}
	err = sub.BindQueue(queue, exchange)
	if err != nil {
		t.Fatal(err)
	}

	size := 1000
	comparisonLookup := make([]int, 0, size)
	for i := 0; i < size; i++ {
		if err := pub.Publish(exchange, "", i); err != nil {
			t.Fatal(err)
		}
		comparisonLookup = append(comparisonLookup, i)
	}

	c, err := sub.Consume(queue)
	if err != nil {
		t.Fatal(err)
	}

	for _, expected := range comparisonLookup {
		select {
		case <-time.After(time.Second * 5):
			t.Fatal("Timeouted after 5 seconds")
		case msg, ok := <-c:
			if !ok {
				// reconnect on unexpected channe closure
				c, err = sub.Consume(queue)
				if err != nil {
					t.Fatal(err)
				}
			}
			got := string(msg.Body)

			if toString(expected) != got {
				t.Fatalf("Expected '%d' != got '%s'", expected, got)
			} else {
				t.Logf("Got: %s, Expected: %s", got, toString(expected))
			}
		}
	}

	broadcast := "broadcast"
	err = sub.CreateExchange(broadcast)
	if err != nil {
		t.Fatal(err)
	}

	// also get messages from the second exchange
	err = sub.BindQueue(queue, broadcast)
	if err != nil {
		t.Fatal(err)
	}

	expectedTest := make([]string, 0)
	expectedBroadcast := make([]string, 0)
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			// as we did use test above already, we cannot reuse the "test" queue here!!!
			expectedTest = append(expectedTest, toString(i))
			if err := pub.Publish(exchange, "", i); err != nil {
				t.Fatal(err)
			}
		} else {
			expectedBroadcast = append(expectedBroadcast, toString(i))
			if err := pub.Publish(broadcast, "", i); err != nil {
				t.Fatal(err)
			}
		}
	}

	for i := 0; i < size; i++ {
		select {
		case <-time.After(time.Second * 5):
			t.Fatal("Timeouted after 5 seconds #2")
		case msg, ok := <-c:
			if !ok {
				// reconnect on unexpected channel closure
				c, err = sub.Consume(queue)
				if err != nil {
					t.Fatal(err)
				}
			}
			got := string(msg.Body)

			if len(expectedTest) > 0 && expectedTest[0] == got {
				t.Logf("Got message '%s' from test queue", got)
				expectedTest = expectedTest[1:]
			} else if len(expectedBroadcast) > 0 && expectedBroadcast[0] == got {
				t.Logf("Got message '%s' from broadcast queue", got)
				expectedBroadcast = expectedBroadcast[1:]
			} else {
				t.Fatal("Unexpected state, message is neither from test, nor from broadcast queue")
			}
		}
	}

}
