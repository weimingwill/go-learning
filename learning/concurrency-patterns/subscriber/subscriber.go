package subscriber

type Subscriber interface {
	Notify(interface{}) error
	Close()
}
