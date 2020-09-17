package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/internal/web/apiclient"
)

func New(client *apiclient.Client) *Server {
	return &Server{
		client: client,
	}
}

type Server struct {
	client *apiclient.Client
}

func (s *Server) CreateRoutes(m *http.ServeMux) {
	r := mux.NewRouter()
	r.HandleFunc("/", s.index)
	r.HandleFunc("/new", s.newGame)
	r.HandleFunc("/{game}", s.gameview)
	r.HandleFunc("/{game}/play", s.play)
	m.Handle("/", r)
}
