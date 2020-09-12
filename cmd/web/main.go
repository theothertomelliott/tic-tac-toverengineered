package main

import (
	"log"
	"net/http"
	"os"

	"github.com/theothertomelliott/tic-tac-toverengineered/internal/web"
	"github.com/theothertomelliott/tic-tac-toverengineered/internal/web/apiclient"
)

func getAPIBaseURL() string {
	if apiBaseURL := os.Getenv("API_BASE_URL"); apiBaseURL != "" {
		return apiBaseURL
	}
	return "http://localhost:8081"
}

func main() {
	log.Println("Starting web")
	server := web.New(apiclient.New(getAPIBaseURL()))

	mux := http.NewServeMux()
	server.CreateRoutes(mux)

	log.Println("Listening on port :8080")
	http.ListenAndServe(":8080", mux)
}
