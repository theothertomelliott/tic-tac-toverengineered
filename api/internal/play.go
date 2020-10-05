package api

import (
	"net/http"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/param"
)

func (s *Server) playHandler(w http.ResponseWriter, req *http.Request) {
	gameID, err := s.verifyID(w, req)
	if err != nil {
		return
	}

	var player player.Mark
	if err := param.Parse(req, "player", &player, param.ParseOptions{Required: true}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var pos grid.Position
	if err := param.Parse(req, "pos", &pos, param.ParseOptions{Required: true}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.turn.TakeTurn(req.Context(), gameID, player, pos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(w, "ok")
}
