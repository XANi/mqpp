package backend

import (
	"fmt"
	"github.com/XANi/mqpp/backend/amqp"
	"github.com/XANi/mqpp/backend/mqtt"
	"github.com/XANi/mqpp/common"
)

func Connect(mq string, url string, opts interface{}) (common.Backend, error) {
	switch mq {
	case "amqp":
		return amqp.New(url, opts)
	case "mqtt":
		return mqtt.New(url, opts)
	default:
		return nil, fmt.Errorf("no such backend: %s", mq)

	}
	return nil, nil
}
