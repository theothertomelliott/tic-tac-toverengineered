package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/api/pkg/tictactoeapi/tictactoeapiclient"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func New(apiclient *tictactoeapiclient.Client) *Server {
	return &Server{
		apiclient: apiclient,
	}
}

type Server struct {
	apiclient *tictactoeapiclient.Client
}

func (s *Server) AddRoutes(r *mux.Router) {
	fmt.Println("Adding routes")
	r.Handle("/", otelhttp.NewHandler(http.HandlerFunc(s.index), "index"))
	r.Handle("/new", otelhttp.NewHandler(http.HandlerFunc(s.newGame), "newgame"))
	r.Handle("/{game}", otelhttp.NewHandler(http.HandlerFunc(s.gameview), "gameview"))
	r.Handle("/{game}/play", otelhttp.NewHandler(http.HandlerFunc(s.play), "play"))
}
