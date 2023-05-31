package dr

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DR struct {
	config          *Config
	db              *sql.DB
	orderRepository *OrderRepository
}

func New(config *Config) *DR {
	return &DR{
		config: config,
	}
}

func (d *DR) Open() error {
	db, err := sql.Open("postgres", d.config.DatabaseURL)
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных:", err)
		return err
	}

	d.db = db
	return nil
}

func (d *DR) Close() {
	d.db.Close()
}

func (d *DR) Order() *OrderRepository {
	if d.orderRepository != nil {
		return d.orderRepository
	}
	d.orderRepository = &OrderRepository{
		dr: d,
	}
	return d.orderRepository
}
