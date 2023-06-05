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
			s.Writer(w, []byte("{message: failed to parse form data}"), http.StatusBadRequest)
			return
		}

		order := []byte(r.Form.Get("order"))
		if len(order) == 0 {
			s.Writer(w, []byte("{message: the order variable in the form is empty}"), http.StatusBadRequest)
			return
		}

		err := s.broker.Publish(strconv.Itoa(messager.OrderCreateOrder), order)
		if err != nil {
			return
		}

		s.Writer(w, []byte("{message: order created successfully}"), http.StatusOK)
	}
}

func (s *server) getOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/get-order/"))
		if err != nil {
			s.Writer(w, []byte("{message: id is not integer}"), http.StatusBadRequest)
			return
		}

		orderId := fmt.Sprintf("order_id_%d", id)
		order, err := s.cache.Get(orderId)
		if err != nil {
			value, err := s.storage.Order().Get(id)
			if err != nil {
				s.Writer(w, []byte("{message: "+err.Error()+"}"), http.StatusNotFound)
				return
			}
			s.cache.Set(orderId, value, 0)
			order, err = s.cache.Get(orderId)
		}

		s.Writer(w, []byte(order), http.StatusOK)
	}
}

func (s *server) getAllOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		orderIds, err := s.storage.Order().GetAllOrdersId()
		if err != nil {
			s.Writer(w, []byte("{message: "+err.Error()+"}"), http.StatusNotFound)
			return
		}
		s.Writer(w, []byte(orderIds), http.StatusOK)
	}
}

func (s *server) Writer(w http.ResponseWriter, message []byte, httpStatus int) {
	w.Write(message)
	w.WriteHeader(httpStatus)
}
