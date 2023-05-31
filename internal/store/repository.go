package store

import "github.com/ent1k1377/wb_l0/internal/model"

type OrderRepository interface {
	Create(order *model.Order) error
}
