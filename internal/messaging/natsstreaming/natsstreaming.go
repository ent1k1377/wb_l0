package natsstreaming

import (
	"github.com/nats-io/stan.go"
)

type Stan struct {
	Conn stan.Conn
}

func New(conn stan.Conn) *Stan {
	return &Stan{
		Conn: conn,
	}
}

func (s *Stan) Publish(channel string, data []byte) error {
	return s.Conn.Publish(channel, data)
}

func (s *Stan) Subscribe(channel string, cb func(data []byte)) (unsub func() error, err error) {
	sub, err := s.Conn.Subscribe(channel, func(msg *stan.Msg) {
		data := msg.Data
		cb(data)
	})
	if err != nil {
		return nil, err
	}
	return sub.Unsubscribe, nil
}
