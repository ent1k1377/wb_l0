package storage

import "github.com/ent1k1377/wb_l0/internal/model"

type OrderRepository interface {
	Create(order *model.Order) error
	Get(id int) (string, error)
}
