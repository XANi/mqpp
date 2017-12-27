package main

import (
	"github.com/XANi/mqpp/backend"
	"github.com/XANi/mqpp/common"
	"github.com/urfave/cli"
	"os"
	"sync"
)

type Connections struct {
	conns map[string]common.Backend
	sync.Mutex
}

var mqConnections Connections

func init() {
	mqConnections.conns = make(map[string]common.Backend)
}

func Get(c *cli.Context) error {
	allURLDefault := true
	for k, v := range defaultUrls {
		if c.GlobalString(k+"-url") != v {
			allURLDefault = false
			break
		}
	}
	if allURLDefault && (c.GlobalString("mq-type") == "") {
		log.Notice("All queue URLs are default, will try to connect to each one in turn")
	}
	if c.GlobalString("mq-type") == "" {
		var wg sync.WaitGroup
		for _, mq := range supportedMQ {
			// supress errors when connecting to queues using default URLs
			supressError := false
			url := defaultUrls[mq]
			if c.GlobalString(mq+"-url") == defaultUrls[mq] {
				supressError = true
			} else {
				url = c.GlobalString(mq + "-url")
			}
			wg.Add(1)
			go func(mq string, url string) {
				defer wg.Done()
				conn, err := backend.Connect(mq, url, nil)
				if err == nil {
					log.Noticef("connected to %s:%s", mq, url)
					mqConnections.Lock()
					mqConnections.conns[mq] = conn
					mqConnections.Unlock()
				} else {
					if !supressError || true {
						log.Warningf("connection to %s failed: %s", url, err)
					}
				}
			}(mq, url)
		}
		wg.Wait()
		if len(mqConnections.conns) < 1 {
			log.Errorf("Nothing connected, exiting")
			os.Exit(1)
		}
	}
	for k, v := range mqConnections.conns {
		go func(k string, v common.Backend) {
			ch := v.GetDefault()
			for msg := range ch {
				log.Infof("%s: %s", k, msg)
			}
			log.Warningf("connector %s closed the connection")
		}(k, v)
	}
	return nil
}