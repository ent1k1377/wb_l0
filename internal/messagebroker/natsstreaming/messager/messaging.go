package messager

import (
	"github.com/ent1k1377/wb_l0/internal/cache"
	"github.com/ent1k1377/wb_l0/internal/messagebroker/natsstreaming"
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
	cache     cache.Cache
	stan      *natsstreaming.Stan
	messaging []Messaging
}

func StartListeningPublisher(store storage.Storage, cache cache.Cache, stan *natsstreaming.Stan) {
	message := New(store, cache, stan).messaging
	for _, m := range message {
		m.StartListening()
	}
}

func New(store storage.Storage, cache cache.Cache, stan *natsstreaming.Stan) *Message {
	return &Message{
		store:     store,
		cache:     cache,
		stan:      stan,
		messaging: initMessages(store, cache, stan),
	}
}

func initMessages(store storage.Storage, cache cache.Cache, stan *natsstreaming.Stan) []Messaging {
	messages := []Messaging{
		&OrderMessaging{store: store, cache: cache, stan: stan},
	}
	return messages
}
