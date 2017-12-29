package common

// Message is common message type used by printing/formatting function. it is backend job to fit incoming messages into that template
type Message struct {
	// Source split by path separator
	Source  []string
	Headers map[string]interface{}
	Body    []byte
}
