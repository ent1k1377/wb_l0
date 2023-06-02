package sqlstorage

import (
	"database/sql"
	"github.com/ent1k1377/wb_l0/internal/cache"
	"github.com/ent1k1377/wb_l0/internal/storage"
)

type Storage struct {
	db              *sql.DB
	cache           cache.Cache
	orderRepository *OrderRepository
}

func New(db *sql.DB, cache cache.Cache) *Storage {
	return &Storage{
		db:    db,
		cache: cache,
	}
}

func (s *Storage) Order() storage.OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}
	s.orderRepository = &OrderRepository{
		storage: s,
	}
	return s.orderRepository
}
