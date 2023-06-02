package apiserver

import (
	"fmt"
	"github.com/ent1k1377/wb_l0/internal/cache"
	"github.com/ent1k1377/wb_l0/internal/messagebroker"
	"github.com/ent1k1377/wb_l0/internal/messagebroker/natsstreaming/messager"
	"github.com/ent1k1377/wb_l0/internal/storage"
	"net/http"
	"strconv"
	"strings"
)

type server struct {
	router  *http.ServeMux
	storage storage.Storage
	broker  messagebroker.Broker
	cache   cache.Cache
}

func newServer(storage storage.Storage, broker messagebroker.Broker, cache cache.Cache) *server {
	s := &server{
		router:  http.NewServeMux(),
		storage: storage,
		broker:  broker,
		cache:   cache,
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/create-order/", s.createOrder())
	s.router.HandleFunc("/get-order/", s.getOrder())
	s.router.HandleFunc("/get-all-orders/", s.getAllOrders())
}

func (s *server) createOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}

		if r.ParseForm() != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		jsonData := []byte(r.Form.Get("order"))
		err := s.broker.Publish(strconv.Itoa(messager.OrderCreateOrder), jsonData)
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *server) getOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/get-order/"))
		if err != nil {
			http.Error(w, "id is not integer", http.StatusBadRequest)
			return
		}

		orderId := fmt.Sprintf("order_id_%d", id)
		valueC, err := s.cache.Get(orderId)
		if err != nil {
			value, err := s.storage.Order().Get(id)
			if err != nil {
				w.Write([]byte("{response: " + err.Error() + "}"))
				w.WriteHeader(http.StatusNotFound)
				return
			}
			s.cache.Set(orderId, value, 0)

			valueC, err = s.cache.Get(orderId)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(valueC))
	}
}

func (s *server) getAllOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{response: jopa}"))
	}
}
