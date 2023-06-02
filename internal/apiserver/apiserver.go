package apiserver

import (
	"database/sql"
	"github.com/ent1k1377/wb_l0/internal/cache/rediscache"
	"github.com/ent1k1377/wb_l0/internal/messagebroker/natsstreaming"
	"github.com/ent1k1377/wb_l0/internal/messagebroker/natsstreaming/messager"
	"github.com/ent1k1377/wb_l0/internal/storage/sqlstorage"
	"github.com/go-redis/redis"
	"github.com/nats-io/stan.go"
	"net/http"
	"os"
)

func Start(config *Config) error {
	// Подключение к бд
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	// Подключение к redis
	client, err := newRedis(config.RedisURL)
	if err != nil {
		return err
	}

	// Подключение к nats-streaming
	conn, err := newStan(config.StanURL)
	if err != nil {
		return err
	}

	// Инициализация сервисов и запуск сервера
	cache := rediscache.New(client)
	storage := sqlstorage.New(db, cache)
	broker := natsstreaming.New(*conn)
	server := newServer(storage, broker, cache)

	// Прослушивание publisher
	messager.StartListeningPublisher(storage, cache, broker)

	return http.ListenAndServe(config.BindAddr, server)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open(os.Getenv("DB"), dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func newStan(natsURL string) (*stan.Conn, error) {
	sc, err := stan.Connect(os.Getenv("STAN_CLUSTER_ID"), "client-0", stan.NatsURL(natsURL))
	if err != nil {
		return nil, err
	}
	return &sc, nil
}

func newRedis(redisURL string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisURL,
		DB:   0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
