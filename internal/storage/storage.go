package storage

import (
	_ "github.com/lib/pq"
)

type Storage interface {
	Order() OrderRepository
}
