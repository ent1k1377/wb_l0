package store

import (
	_ "github.com/lib/pq"
)

type Store interface {
	Order() OrderRepository
}
