package common

import "time"

// Message is common message type used by printing/formatting function. it is backend job to fit incoming messages into that template
type Message struct {
	// Source split by path separator
	TS time.Time
	// true if it is TS from actual queue, false if it is just receive time
	TSReliable bool
	Source  []string
	Headers map[string]interface{}
	Body    []byte
}
