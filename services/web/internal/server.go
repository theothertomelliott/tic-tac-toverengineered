package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi/tictactoeapiclient"
)

func New(apiclient *tictactoeapiclient.Client) *Server {
	return &Server{
		apiclient: apiclient,
	}
}

type Server struct {
	apiclient *tictactoeapiclient.Client
}

// wrap will wrap an http handler with all intended middleware
func wrap(handler http.HandlerFunc, name string) http.Handler {
	return monitoring.WrapHTTP(handler, name)
}

func (s *Server) AddRoutes(r *mux.Router) {
	fmt.Println("Adding routes")
	r.Handle("/", wrap(s.index, "index"))
	r.Handle("/new", wrap(s.newGame, "newgame"))
	r.Handle("/{game}", wrap(s.gameview, "gameview"))
	r.Handle("/{game}/play", wrap(s.play, "play"))
}
