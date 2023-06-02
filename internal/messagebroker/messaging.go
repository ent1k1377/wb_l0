package messagebroker

type Broker interface {
	Publish(topic string, data []byte) error
	Subscribe(topic string, cb func(data []byte)) (unsub func() error, err error)
}
