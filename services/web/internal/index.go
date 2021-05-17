package web

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/theothertomelliott/tic-tac-toverengineered/common/http/param"
	"github.com/theothertomelliott/tic-tac-toverengineered/services/gamerepo/pkg/game"
)

var indexTmpl string

func init() {
	content, err := ioutil.ReadFile("views/index.html")
	if err != nil {
		panic(err)
	}
	indexTmpl = string(content)
}

func (s *Server) index(w http.ResponseWriter, req *http.Request) {
	var max, offset int32
	if err := param.Parse(req, "max", &max, param.ParseOptions{Default: 10}); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	if err := param.Parse(req, "offset", &offset, param.ParseOptions{Default: 0}); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	t, err := template.New("webpage").Parse(indexTmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Games      []game.ID
		PrevOffset int32
		NextOffset int32
		Max        int32
	}{
		PrevOffset: offset - max,
		NextOffset: offset + max,
		Max:        max,
	}
	if data.PrevOffset < 0 {
		data.PrevOffset = 0
	}

	if err := s.client.RawApiGet(req.Context(), fmt.Sprintf("/?max=%v&offset=%v", max, offset), &data.Games); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
