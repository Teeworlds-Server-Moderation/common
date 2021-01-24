package mqtt

import (
	"log"
	"strconv"
)

// Message is a simple struct that allows to publish to different topics
// with a single publisher client connection
type Message struct {
	Topic   string
	Payload string
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
	default:
		log.Panicf("Invalid type passed: %T", t)
	}
	return ""
}
