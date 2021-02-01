package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/theothertomelliott/tic-tac-toverengineered/api/pkg/apiclient"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring/defaultmonitoring"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/version"
	web "github.com/theothertomelliott/tic-tac-toverengineered/web/internal"
)

func getAPIBaseURL() string {
	if apiBaseURL := os.Getenv("API_BASE_URL"); apiBaseURL != "" {
		return apiBaseURL
	}
	return "http://localhost:8081"
}

func main() {
	version.Println()
	defaultmonitoring.Init("web")
	defer monitoring.Close()

	log.Println("Starting web")
	client := &http.Client{
		Transport: monitoring.WrapHTTPTransport(http.DefaultTransport),
		Timeout:   time.Second * 5,
	}
	server := web.New(apiclient.New(getAPIBaseURL(), client))

	log.Println("Listening on port :8080")

	m := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public"))
	m.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	prefix := os.Getenv("ROUTE_PREFIX")
	if prefix != "" && prefix != "/" {
		routes := m.PathPrefix(prefix).Subrouter()
		server.AddRoutes(routes)
	} else {
		server.AddRoutes(m)
	}

	http.ListenAndServe(":8080", monitoring.WrapHTTP(m))
}

type prefixHandler struct {
	prefix string
}

func (p *prefixHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, p.prefix)
	fmt.Fprint(w, r.URL.Path)
}
