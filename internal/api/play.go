package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

func (s *Server) playHandler(w http.ResponseWriter, req *http.Request) {
	gameID := game.ID(mux.Vars(req)["game"])
	playerParams, ok := req.URL.Query()["player"]
	if !ok || len(playerParams) == 0 {
		http.Error(w, "player is required", http.StatusInternalServerError)
		return
	}
	player := player.Mark(playerParams[0])

	posParams, ok := req.URL.Query()["pos"]
	if !ok || len(posParams) == 0 {
		http.Error(w, "pos is required", http.StatusInternalServerError)
		return
	}

	var pos grid.Position
	if err := json.Unmarshal([]byte(posParams[0]), &pos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.turn.TakeTurn(gameID, player, pos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(w, "ok")
}
