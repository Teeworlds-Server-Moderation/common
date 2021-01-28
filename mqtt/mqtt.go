package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

// this package can be used to define reading and writing routines that connect to the Mosquitto broker
// those routines are supposed to run in their own goroutines and either receive or pass events through channels
const (
	QOS = 1
)

var (
	// Debug for debug log messages
	Debug = false
)

var (
	DefaultPublishHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("unexpected message: %s\n", msg)
	}

	DefaultOnConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Println("connection lost")
	}

	DefaultOnReconnectingHandler = func(client mqtt.Client, option *mqtt.ClientOptions) {
		fmt.Println("attempting to reconnect")
	}
)

// uniqueClientID creates a client ID that is unique and also not longer than 23 characters
// as per MQTT specification
func uniqueClientID(prefix string) string {
	prefix = prefix + "-"
	id := uuid.New().String()
	unique := prefix + id[len(id)-23+len(prefix):]
	return unique
}
