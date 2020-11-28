package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/theothertomelliott/tic-tac-toverengineered/api/pkg/apiclient"
	"github.com/theothertomelliott/tic-tac-toverengineered/common/monitoring"
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

	log.Println("Starting web")
	client := &http.Client{
		Transport: monitoring.WrapHTTPTransport(http.DefaultTransport),
		Timeout:   time.Second * 5,
	}
	server := web.New(apiclient.New(getAPIBaseURL(), client))

	mux := http.NewServeMux()
	server.CreateRoutes(mux)

	closeMonitoring := monitoring.Init("web")
	defer closeMonitoring()

	log.Println("Listening on port :8080")

	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/", http.StripPrefix("/public/", fs))

	http.ListenAndServe(":8080", monitoring.WrapHTTP(mux))
}
