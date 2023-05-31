package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/ent1k1377/wb_l0/internal/dr"
	"github.com/ent1k1377/wb_l0/internal/model"
	"github.com/nats-io/stan.go"
	"net/http"
	"os"
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
	http.Handle("/create_order/", createOrder(s))
}

func createOrder(s *APIServer) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}

		fmt.Println(r.Body)
		jsonData := []byte("{\n  \"order_uid\": \"b563feb7b2b84b6test\",\n  \"track_number\": \"WBILMTESTTRACK\",\n  \"entry\": \"WBIL\",\n  \"delivery\": {\n    \"name\": \"Test Testov\",\n    \"phone\": \"+9720000000\",\n    \"zip\": \"2639809\",\n    \"city\": \"Kiryat Mozkin\",\n    \"address\": \"Ploshad Mira 15\",\n    \"region\": \"Kraiot\",\n    \"email\": \"test@gmail.com\"\n  },\n  \"payment\": {\n    \"transaction\": \"b563feb7b2b84b6test\",\n    \"request_id\": \"\",\n    \"currency\": \"USD\",\n    \"provider\": \"wbpay\",\n    \"amount\": 1817,\n    \"payment_dt\": 1637907727,\n    \"bank\": \"alpha\",\n    \"delivery_cost\": 1500,\n    \"goods_total\": 317,\n    \"custom_fee\": 0\n  },\n  \"items\": [\n    {\n      \"chrt_id\": 9934930,\n      \"track_number\": \"WBILMTESTTRACK\",\n      \"price\": 453,\n      \"rid\": \"ab4219087a764ae0btest\",\n      \"name\": \"Mascaras\",\n      \"sale\": 30,\n      \"size\": \"0\",\n      \"total_price\": 317,\n      \"nm_id\": 2389212,\n      \"brand\": \"Vivienne Sabo\",\n      \"status\": 202\n    }\n  ],\n  \"locale\": \"en\",\n  \"internal_signature\": \"\",\n  \"customer_id\": \"test\",\n  \"delivery_service\": \"meest\",\n  \"shardkey\": \"9\",\n  \"sm_id\": 99,\n  \"date_created\": \"2021-11-26T06:22:19Z\",\n  \"oof_shard\": \"1\"\n}")
		var orderData model.Order
		err := json.Unmarshal(jsonData, &orderData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = s.dr.Order().Create(&orderData)
		if err != nil {
			fmt.Println(err)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, orderData)
	}
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
