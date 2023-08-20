package mqtt

import (
	"fmt"
	"github.com/XANi/mqpp/mq"
	pmqtt "github.com/eclipse/paho.mqtt.golang"
	"net/url"
)

type MQTT struct {
	cfg  Config
	conn pmqtt.Client
}

type Config struct {
	Filter string
}

func New(urlAddr string, cfg Config) (*MQTT, error) {
	var backend MQTT
	urlParsed, err := url.Parse(urlAddr)
	if err != nil {
		return nil, fmt.Errorf("Can't parse url [%s]:%s", urlAddr, err)
	}
	clientOpts := pmqtt.NewClientOptions().AddBroker(urlAddr)
	if urlParsed.User != nil && urlParsed.User.Username() != "" {
		clientOpts.Username = urlParsed.User.Username()
		clientOpts.Password, _ = urlParsed.User.Password()
	}

	if len(cfg.Filter) == 0 {
		cfg.Filter = "#"
	}
	backend.cfg = cfg
	client := pmqtt.NewClient(clientOpts)
	if connectToken := client.Connect(); connectToken.Wait() && connectToken.Error() != nil {
		return nil, fmt.Errorf("Could not connect to MQTT: %s", connectToken.Error())
	}
	backend.conn = client
	return &backend, nil
}

func (m *MQTT) Subscribe(topic string) chan mq.Event {
	ch := make(chan mq.Event, 1)
	token := m.conn.Subscribe(topic, 0, func(cl pmqtt.Client, msg pmqtt.Message) {
		ev := mq.Event{
			Topic:   msg.Topic(),
			Payload: msg.Payload(),
		}
		ch <- ev
	})
	_ = token // todo handle it somehow
	return ch
}
