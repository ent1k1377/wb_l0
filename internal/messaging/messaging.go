package messaging

type Messenger interface {
	Publish(topic string, data []byte) error
	Subscribe(topic string, cb func(data []byte)) (unsub func() error, err error)
}
