package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

func (s *Server) playHandler(w http.ResponseWriter, req *http.Request) {
	gameID := game.ID(mux.Vars(req)["game"])

	var player player.Mark
	if err := parseParam(req, "player", &player); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var pos grid.Position
	if err := parseParam(req, "pos", &pos); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.turn.TakeTurn(gameID, player, pos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(w, "ok")
}
