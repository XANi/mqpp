package common

type Backend interface {
	//Connect(url string, opts interface{}) error
	Get()
	GetDefault() chan Message
}
