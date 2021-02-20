package amqp

import (
	"encoding/json"
	"log"
	"strconv"
)

// Message is a simple struct that allows to publish to different topics
// with a single publisher client connection
type Message struct {
	Exchange string
	Queue    string
	Payload  string
}

func toString(i interface{}) string {
	switch t := i.(type) {
	case string:
		return t
	case []byte:
		return string(t)
	case int:
		return strconv.Itoa(t)
	case int64:
		return strconv.FormatInt(t, 10)
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 32)
	case bool:
		return strconv.FormatBool(t)
	case Message:
		return t.Payload
	}
	b, err := json.Marshal(i)
	if err != nil {
		log.Panicf("Invalid type passed: %T", i)
	}
	return string(b)
}
