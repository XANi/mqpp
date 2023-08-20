package tui

import (
	"github.com/XANi/mqpp/mq"
	"github.com/pterm/pterm"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Config struct {
	MQChannel chan mq.Event
	Logger    *zap.SugaredLogger
}

type EvTree struct {
	key         string
	t           map[string]*EvTree
	lastMessage []byte
	ts          time.Time
}

type TUI struct {
	mq     chan mq.Event
	evTree *EvTree
	l      *zap.SugaredLogger
}

func New(cfg Config) (*TUI, error) {
	t := TUI{
		mq: cfg.MQChannel,
		evTree: &EvTree{
			t:           make(map[string]*EvTree),
			lastMessage: nil,
			ts:          time.Time{},
		},
		l: cfg.Logger,
	}
	return &t, nil
}

func (t *TUI) updateTree(ev mq.Event) {
	//defer func() {
	//	if r := recover(); r != nil {
	//		t.l.Infof(pp.Sprint(t.evTree))
	//		t.l.Infof(pp.Sprint(ev))
	//		t.l.Errorf("recover: %s", r)
	//		os.Exit(1)
	//	}
	//}()
	path := strings.Split(ev.Topic, "/")
	pt := t.evTree
	last := false
	for idx, el := range path {
		if idx == len(path)-1 {
			last = true
		}
		if pt == nil {
			pt = &EvTree{
				t: make(map[string]*EvTree, 0),
			}
		}

		if v, ok := pt.t[el]; ok {
			pt = v
		} else {
			pt.t[el] = &EvTree{
				t:   make(map[string]*EvTree, 0),
				key: el,
			}
		}
		pt = pt.t[el]
		if last {
			pt.ts = time.Now()
			pt.lastMessage = ev.Payload
			pt.t = make(map[string]*EvTree, 0)
			pt.key = el
			return
		}
	}
}

func (t *TUI) Run() error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	a := pterm.DefaultArea
	a.Fullscreen = true
	area, _ := a.Start()

	for {
		timeout := time.After(time.Second)
		select {
		case <-timeout:
		case <-sigs:
			os.Exit(0)
		case e := <-t.mq:
			t.updateTree(e)
		}
		time.Sleep(time.Second)
		s, err := t.EventTree()
		if err != nil {
			return err
		}
		area.Update(s)
	}

	return nil

}
