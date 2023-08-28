package shared

import (
	"log"
	"math/rand"
)

// MQTT
const MQTTHost = "tcp://localhost:1883"
const MQTTTopic = "PubSub"
const MQTTRequest = "request"
const MQTTReply = "reply"

// Other configurations
const SampleSize = 10000
const RequestQueue = "request_queue"
const ResponseQueue = "response_queue"

type Request struct {
	ClientID string
	Keywords string
}

type Reply struct {
	Books string
}

func ChecaErro(err error, msg string) {
	if err != nil {
		log.Fatalf("%s!!: %s", msg, err)
	}
	//fmt.Println(msg)
}

func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
