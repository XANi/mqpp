package backend

import (
	"fmt"
	"github.com/XANi/mqpp/backend/amqp"
	"github.com/XANi/mqpp/backend/mqtt"
	"github.com/XANi/mqpp/common"
	"strings"
)

func Connect(mq string, url string, topicFilter string) (common.Backend, error) {
	switch mq {
	case "amqp":
		f := strings.Replace(topicFilter,"/",".",-1)
		return amqp.New(url, amqp.AMQPConfig{
			Filter: f,
		})
	case "mqtt":
		return mqtt.New(url, mqtt.MQTTConfig{
			Filter: topicFilter,
		})
	default:
		return nil, fmt.Errorf("no such backend: %s", mq)

	}
	return nil, nil
}
