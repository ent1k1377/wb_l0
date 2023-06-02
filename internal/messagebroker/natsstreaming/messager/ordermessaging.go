package messager

import (
	"encoding/json"
	"fmt"
	"github.com/ent1k1377/wb_l0/internal/cache"
	"github.com/ent1k1377/wb_l0/internal/messagebroker/natsstreaming"
	"github.com/ent1k1377/wb_l0/internal/model"
	"github.com/ent1k1377/wb_l0/internal/storage"
	"log"
	"strconv"
)

type OrderMessaging struct {
	store storage.Storage
	cache cache.Cache
	stan  *natsstreaming.Stan
}

func (o *OrderMessaging) StartListening() {
	go o.CreateOrder()
}

func (o *OrderMessaging) CreateOrder() {
	unsubscription, err := o.stan.Subscribe(strconv.Itoa(OrderCreateOrder), o.handleOrderMessage)
	if err != nil {
		log.Fatalf("Failed to subscribe to order messages: %v", err)
	}

	defer unsubscription()

	select {}
}

func (o *OrderMessaging) handleOrderMessage(data []byte) {
	var orderData model.Order
	err := json.Unmarshal(data, &orderData)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = o.store.Order().Create(&orderData)

	orderId := fmt.Sprintf("order_id_%s", orderData.OrderUID)
	fmt.Println(orderId)
	err = o.cache.Set(orderId, string(data), 0)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}
}
