package apiserver

import (
	"github.com/ent1k1377/wb_l0/internal/messaging"
	"github.com/ent1k1377/wb_l0/internal/messaging/natsstreaming/messager"
	"github.com/ent1k1377/wb_l0/internal/store"
	"net/http"
	"strconv"
)

type server struct {
	router    *http.ServeMux
	store     store.Store
	messenger messaging.Messenger
}

func newServer(store store.Store, messenger messaging.Messenger) *server {
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
	s.router.HandleFunc("/create_order", s.createOrder())
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
