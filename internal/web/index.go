package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/param"
)

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
		Games []game.ID
	}{}

	if err := s.client.RawApiGet(fmt.Sprintf("/?max=%v&offset=%v", max, offset), &data.Games); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

const indexTmpl = `
<html>
<head>
<title>Tic Tac Toe</title>
</head>
<body>
	<h1>Tic Tac Toe</h1>
	<p><a href="/new">New Game</a></p>
	<ul>
	{{ range .Games}}
		<li><a href="/{{ . }}">{{ . }}</a></li>
	{{ end }}
	</ul>
</body>
</html>
`
