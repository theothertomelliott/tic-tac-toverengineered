package web

import (
	"html/template"
	"net/http"

	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/game"
)

func (s *Server) index(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("webpage").Parse(indexTmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Games []game.ID
	}{}

	if err := s.client.RawApiGet("/?max=10&offset=0", &data.Games); err != nil {
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
