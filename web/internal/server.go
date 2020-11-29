package web

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/api/pkg/apiclient"
)

func New(client *apiclient.Client) *Server {
	return &Server{
		client: client,
	}
}

type Server struct {
	client *apiclient.Client
}

func (s *Server) AddRoutes(r *mux.Router) {
	fmt.Println("Adding routes")
	r.HandleFunc("/", s.index)
	r.HandleFunc("/new", s.newGame)
	r.HandleFunc("/{game}", s.gameview)
	r.HandleFunc("/{game}/play", s.play)
}
