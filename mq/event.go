package mq

type Event struct {
	Topic   string
	Payload []byte
}
