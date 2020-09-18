package api

import (
	"net/http"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/grid"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/player"
)

func (s *Server) gridHandler(w http.ResponseWriter, req *http.Request) {
	gameID, err := s.verifyID(w, req)
	if err != nil {
		return
	}

	var out [][]*player.Mark
	for i := 0; i < 3; i++ {
		var row []*player.Mark
		for j := 0; j < 3; j++ {
			m, err := s.grid.Mark(gameID, grid.Position{X: i, Y: j})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			row = append(row, m)
		}
		out = append(out, row)
	}
	jsonResponse(w, out)
}
