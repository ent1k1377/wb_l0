package apiserver

import (
	"fmt"
	"github.com/ent1k1377/wb_l0/internal/dr"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/nats-io/stan.go"
)

type APIServer struct {
	config *Config
	dr     *dr.DR
	stan   *stan.Conn
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()

	if err := s.configureDB(); err != nil {
		return err
	}
	if err := s.configureStan(); err != nil {
		return err
	}

	return http.ListenAndServe(s.config.BindAddr, nil)
}

func (s *APIServer) configureRouter() {
	http.HandleFunc("/order-id/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return
	}

	fmt.Fprintf(w, "order-id: %d", id)
}

func (s *APIServer) configureDB() error {
	dr := dr.New(s.config.DR)
	if err := dr.Open(); err != nil {
		return err
	}
	s.dr = dr
	return nil
}

func (s *APIServer) configureStan() error {
	natsURL := fmt.Sprintf("nats://%s:%s",
		os.Getenv("STAN_CONTAINER_NAME"),
		os.Getenv("STAN_CONTAINER_PORT"),
	)
	sc, err := stan.Connect(os.Getenv("STAN_CLUSTER_ID"), "client-0", stan.NatsURL(natsURL))
	if err != nil {
		return err
	}

	s.stan = &sc
	return nil
}
