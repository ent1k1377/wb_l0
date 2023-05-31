package apiserver

import (
	"database/sql"
	"github.com/ent1k1377/wb_l0/internal/messaging"
	"github.com/ent1k1377/wb_l0/internal/store/sqlstore"
	"github.com/nats-io/stan.go"
	"net/http"
	"os"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	conn, err := newStan(config.StanURL)
	if err != nil {
		return err
	}

	store := sqlstore.New(db)
	stan := messaging.New(conn)
	server := newServer(store, *stan)

	return http.ListenAndServe(config.BindAddr, server)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
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
