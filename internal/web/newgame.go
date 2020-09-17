package web

import (
	"fmt"
	"net/http"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
)

func (s *Server) newGame(w http.ResponseWriter, req *http.Request) {
	var gameID game.ID
	if err := s.client.RawApiGet("/new", &gameID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, req, fmt.Sprintf("/%v", gameID), http.StatusFound)
}
