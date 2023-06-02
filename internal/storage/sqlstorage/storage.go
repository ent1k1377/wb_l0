package sqlstorage

import (
	"database/sql"
	"github.com/ent1k1377/wb_l0/internal/storage"
)

type Storage struct {
	db              *sql.DB
	orderRepository *OrderRepository
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) Order() storage.OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}
	s.orderRepository = &OrderRepository{
		store: s,
	}
	return s.orderRepository
}
