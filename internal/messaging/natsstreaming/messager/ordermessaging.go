package messager

import (
	"encoding/json"
	"fmt"
	"github.com/ent1k1377/wb_l0/internal/messaging/natsstreaming"
	"github.com/ent1k1377/wb_l0/internal/model"
	"github.com/ent1k1377/wb_l0/internal/store"
	"log"
	"strconv"
)

type OrderMessaging struct {
	store store.Store
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
	fmt.Println("create-order")
	var orderData model.Order
	err := json.Unmarshal(data, &orderData)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = o.store.Order().Create(&orderData)
	if err != nil {
		fmt.Println(err)
	}
}
