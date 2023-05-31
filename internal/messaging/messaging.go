package messaging

import (
	"github.com/nats-io/stan.go"
)

type Stan struct {
	conn *stan.Conn
}

func New(conn *stan.Conn) *Stan {
	return &Stan{
		conn: conn,
	}
}
