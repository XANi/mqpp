package mqtt

import (
	backend "github.com/XANi/mqpp/common"
	"github.com/streadway/amqp"
)

type AMQP struct {
	conn *amqp.Connection
}

func New(url string, opts interface{}) (backend.Backend, error) {
	var backend AMQP
	return &backend, nil
}

func (q *AMQP) Get() {

}

func (q *AMQP) GetDefault() chan backend.Message {
	c := make(chan backend.Message, 1)
	go func() {
		c <- backend.Message{Body:[]byte("mqtt dummy")}
	}()
	return c
}
