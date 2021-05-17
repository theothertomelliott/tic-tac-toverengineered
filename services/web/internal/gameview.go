package web

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/player"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/checker/pkg/win"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
)

var gameviewTmpl string

func init() {
	content, err := ioutil.ReadFile("views/game.html")
	if err != nil {
		panic(err)
	}
	gameviewTmpl = string(content)
}

func (s *Server) gameview(w http.ResponseWriter, req *http.Request) {
	gameID := game.ID(mux.Vars(req)["game"])
	t, err := template.New("webpage").Parse(gameviewTmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Game       game.ID
		NextPlayer player.Mark
		Winner     win.Result
		Grid       [][]*player.Mark
	}{
		Game: gameID,
	}
	if err := s.client.ApiGet(req.Context(), gameID, "grid", &data.Grid); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := s.client.ApiGet(req.Context(), gameID, "player/current", &data.NextPlayer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := s.client.ApiGet(req.Context(), gameID, "winner", &data.Winner); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
