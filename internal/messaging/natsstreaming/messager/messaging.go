package messager

import (
	"github.com/ent1k1377/wb_l0/internal/messaging/natsstreaming"
	"github.com/ent1k1377/wb_l0/internal/storage"
)

const (
	OrderCreateOrder = iota
)

type Messaging interface {
	StartListening()
}

type Message struct {
	store     storage.Storage
	stan      *natsstreaming.Stan
	messaging []Messaging
}

func StartListeningPublisher(store storage.Storage, stan *natsstreaming.Stan) {
	messages := New(store, stan).messaging
	for _, m := range messages {
		m.StartListening()
	}
}

func New(store storage.Storage, stan *natsstreaming.Stan) *Message {
	return &Message{
		store:     store,
		stan:      stan,
		messaging: initMessages(store, stan),
	}
}

func initMessages(store storage.Storage, stan *natsstreaming.Stan) []Messaging {
	messages := []Messaging{
		&OrderMessaging{store: store, stan: stan},
	}
	return messages
}
