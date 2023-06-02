package apiserver

import (
	"github.com/ent1k1377/wb_l0/internal/messaging"
	"github.com/ent1k1377/wb_l0/internal/messaging/natsstreaming/messager"
	"github.com/ent1k1377/wb_l0/internal/storage"
	"net/http"
	"strconv"
)

type server struct {
	router    *http.ServeMux
	store     storage.Storage
	messenger messaging.Messenger
}

func newServer(store storage.Storage, messenger messaging.Messenger) *server {
	s := &server{
		router:    http.NewServeMux(),
		store:     store,
		messenger: messenger,
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/create-order", s.createOrder())
	s.router.HandleFunc("/get-all-orders", s.getAllOrders())
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
		err := s.messenger.Publish(strconv.Itoa(messager.OrderCreateOrder), jsonData)
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *server) getAllOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{response: jopa}"))
		w.WriteHeader(http.StatusOK)
	}
}
