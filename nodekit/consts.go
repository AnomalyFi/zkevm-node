package nodekit

import "time"

const (
	publish     = "publish"
	subscribe   = "subscribe"
	unsubscribe = "unsubscribe"
)

const (
	WebsocketEndpoint = "/ws"
)

type Message struct {
	Action  string `json:"action"`
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

const (
	writeWait = 10 * time.Second
)

var (
	BlockTopic = "block"
)

var BlockSubscriptionMessage = Message{
	Action:  subscribe,
	Topic:   BlockTopic,
	Message: "",
}
