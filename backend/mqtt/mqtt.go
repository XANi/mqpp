package mqtt

import (
	"fmt"
    "net/url"
	backend "github.com/XANi/mqpp/common"
    mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
)

type MQTT struct {
	conn mqtt.Client
	cfg *MQTTConfig
}
type MQTTConfig struct{
	Filter string
}

func New(urlAddr string, cfg interface{}) (backend.Backend, error) {
	var backend MQTT
	urlParsed, err := url.Parse(urlAddr)
	if err != nil {
		return nil,fmt.Errorf("Can't parse url [%s]:%s", urlAddr, err)
	}
	clientOpts := mqtt.NewClientOptions().AddBroker(urlAddr)
	if urlParsed.User != nil && urlParsed.User.Username() != "" {
		clientOpts.Username = urlParsed.User.Username()
		clientOpts.Password,_ = urlParsed.User.Password()
	}
	var c MQTTConfig
	if cfg != nil {
		c = cfg.(MQTTConfig)
	}
	if len(c.Filter) == 0 {
		c.Filter = "#"
	}
	backend.cfg = &c
	client := mqtt.NewClient(clientOpts)
	if connectToken := client.Connect(); connectToken.Wait() && connectToken.Error() != nil {
		return nil, fmt.Errorf("Could not connect to MQTT: %s", connectToken.Error())
	}
	backend.conn = client
	return &backend, nil
}

func (q *MQTT) Get() {

}


func (q *MQTT) GetDefault() chan backend.Message {
	c := make(chan backend.Message, 1)
	if token := q.conn.Subscribe(q.cfg.Filter, 0, func(client mqtt.Client, msg mqtt.Message) {
		m := backend.Message{
			Source: strings.Split(msg.Topic(),"/"),
			Body:   msg.Payload(),
		}
		c <- m
	}); token.Wait() && token.Error() != nil {
		c <- backend.Message{Body: []byte(fmt.Sprintf("subscription failed: %s", token.Error()))}
	}
	return c
}
