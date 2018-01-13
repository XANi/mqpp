package main

import (
	"github.com/op/go-logging"
	"github.com/urfave/cli"
	"os"
	"strings"
)

var version string
var log = logging.MustGetLogger("main")

//var stdout_log_format = logging.MustStringFormatter("%{color:bold}%{time:2006-01-02T15:04:05.0000Z-07:00}%{color:reset}%{color} [%{level:.1s}] %{color:reset}%{shortpkg}[%{longfunc}] %{message}")
var stdout_log_format = logging.MustStringFormatter("%{color:bold}[%{level:.1s}] %{color:reset}%{shortpkg}[%{longfunc}] %{message}")

// those need to be mapped to corresponding backend
var defaultUrls = map[string]string{
	"amqp": "amqp://guest:guest@localhost",
	"mqtt": "tcp://127.0.0.1:1883",
}

var supportedMQ = []string{}
var end chan bool

func init() {
	for k := range defaultUrls {
		supportedMQ = append(supportedMQ, k)
	}
}
func main() {
	stderrBackend := logging.NewLogBackend(os.Stderr, "", 0)
	stderrFormatter := logging.NewBackendFormatter(stderrBackend, stdout_log_format)
	logging.SetBackend(stderrFormatter)
	logging.SetFormatter(stdout_log_format)
	app := cli.NewApp()
	app.Name = "mqpp"
	app.Usage = "queue debugger"
	app.Action = func(c *cli.Context) error {
		log.Notice("For usage run with --help")
		Get(c)
		return nil
	}
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "mq-type",
			Usage: "mq type. If not selected will try to connect to all of them. [" + strings.Join(supportedMQ, ", ") + "]",
		},
		cli.StringFlag{
			Name:   "amqp-url",
			Value:  defaultUrls["amqp"],
			Usage:  "AMQP URL (incl. password)",
			EnvVar: "AMQP_URL",
		},
		cli.StringFlag{
			Name:   "mqtt-url",
			Value:  defaultUrls["mqtt"],
			Usage:  "MQTT URL (incl. password)",
			EnvVar: "MQTT_SERVER",
		},
		cli.StringFlag{
			Name:  "topic-filter,filter,topic,t",
			Value: "#",
			Usage: "topic filter for queue",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "get messages from queue/exchange/topic",
			Action:  Get,
		},
	}
	app.Run(os.Args)
}
