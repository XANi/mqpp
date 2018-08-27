package amqp

import (
	"fmt"

	"github.com/XANi/mqpp/common"
	"github.com/streadway/amqp"
	"strings"
	"time"
)

type AMQP struct {
	conn *amqp.Connection
	cfg *AMQPConfig
}
type AMQPConfig struct {
	Filter string
}

func New(url string, cfg interface{}) (common.Backend, error) {
	conn, err := amqp.Dial(url)
	var backend AMQP
	if err != nil {
		return nil, err
	}

	backend.conn = conn
	var c AMQPConfig
	if cfg != nil {
		c = cfg.(AMQPConfig)
	}
	if len(c.Filter) == 0 {
		c.Filter = "#"
	}
	backend.cfg = &c
	return &backend, nil
}

func (q *AMQP) Get() {

}
func (q *AMQP) GetDefault() chan common.Message {
	c := make(chan common.Message, 10)
	go func() {
		ch, err := q.conn.Channel()
		if err != nil {
			c <- common.Message{Body: []byte( fmt.Sprintf("error while getting channel: %s", err))}
			close(c)
			return
		}
		queue, err := ch.QueueDeclare(
			"",    // name of the queue
			false, // durable
			true,  // delete when unused
			true,  // exclusive
			false, // noWait
			nil,   // arguments
		)
		if err != nil {
			c <- common.Message{Body: []byte( fmt.Sprintf("error while declaring queue: %s", err))}
			close(c)
			return
		}
		err = ch.QueueBind(
			queue.Name,  // name of the queue
			q.cfg.Filter,         // bindingKey
			"amq.topic", // sourceExchange
			false,       // noWait
			nil,         // arguments
		)
		if err != nil {
			c <- common.Message{Body: []byte( fmt.Sprintf("error while binding queue: %s", err))}
			close(c)
			return
		}
		events, err := ch.Consume(
			queue.Name, // name
			"mqpp",     // consumerTag,
			false,      // noAck
			false,      // exclusive
			false,      // noLocal
			false,      // noWait
			nil,        // arguments
		)
		c <- common.Message{Body: []byte( "connected to AMQP, consuming messages")}
		for ev := range events {
			var msg common.Message
			source := strings.Split(ev.RoutingKey,".")
			if ev.Timestamp.IsZero() {
				msg.TS = time.Now()
			} else {
				msg.TS = ev.Timestamp
				msg.TSReliable = true
			}
			msg.Source = append([]string{"[" + ev.Exchange + "]"}, source...)
			msg.Body = ev.Body
			msg.Headers = ev.Headers
			c <- msg
		}
		close(c)
	}()

	return c
}
