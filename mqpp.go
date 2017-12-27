package main

import (
	"os"

	"github.com/op/go-logging"
	"github.com/urfave/cli"
)

var version string
var log = logging.MustGetLogger("main")
var stdout_log_format = logging.MustStringFormatter("%{color:bold}%{time:2006-01-02T15:04:05.0000Z-07:00}%{color:reset}%{color} [%{level:.1s}] %{color:reset}%{shortpkg}[%{longfunc}] %{message}")

func main() {
	stderrBackend := logging.NewLogBackend(os.Stderr, "", 0)
	stderrFormatter := logging.NewBackendFormatter(stderrBackend, stdout_log_format)
	logging.SetBackend(stderrFormatter)
	logging.SetFormatter(stdout_log_format)
	app := cli.NewApp()
	app.Name = "mqpp"
	app.Usage = "queue debugger"
	//app.Action = func(c *cli.Context) error {
	//	// if we ever get default action, it goes here
	//	return nil
	//}
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "amqp-url",
			Value:  "amqp://guest:quest@localhost",
			Usage:  "AMQP URL (incl. password)",
			EnvVar: "AMQP_URL",
		},
		cli.StringFlag{
			Name:   "mqtt-url",
			Value:  "mqtt://guest:quest@localhost",
			Usage:  "MQTT URL (incl. password)",
			EnvVar: "MQTT_SERVER",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "get messages from queue/exchange/topic",
			Action: func(c *cli.Context) error {
				os.Exit(0)
				return nil
			},
		},
	}
	app.Run(os.Args)
	log.Error("now add some code!")
}
