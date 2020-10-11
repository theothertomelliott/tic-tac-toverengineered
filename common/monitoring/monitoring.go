package monitoring

import (
	"log"
	"net/http"
	"os"

	"github.com/honeycombio/beeline-go"
	"github.com/honeycombio/beeline-go/wrappers/hnynethttp"
)

// Init sets up monitoring
func Init(serviceName string) func() error {
	key := os.Getenv("HONEYCOMB_API_KEY")
	if len(key) == 0 {
		log.Println("No API key was defined for Honeycomb, telemetry will not be send. Set the HONEYCOMB_API_KEY and HONEYCOMB_DATASET env vars to enable.")
		return func() error { return nil }
	}
	dataset := os.Getenv("HONEYCOMB_DATASET")
	if len(dataset) == 0 {
		log.Println("No dataset name was defined for Honeycomb, telemetry will not be send. Set the HONEYCOMB_API_KEY and HONEYCOMB_DATASET env vars to enable.")
		return func() error { return nil }
	}
	beeline.Init(beeline.Config{
		WriteKey:    key,
		Dataset:     dataset,
		ServiceName: serviceName,
	})
	log.Printf("Initialized Honeycomb for dataset %q", dataset)
	return func() error {
		beeline.Close()
		return nil
	}
}

// WrapHTTP will add monitoring middleware to an http handler
func WrapHTTP(handler http.Handler) http.Handler {
	return hnynethttp.WrapHandler(handler)
}

// WrapHTTPTransport will add monitoring middleware to an HTTP Transport.
// This allows for trace propagation.
func WrapHTTPTransport(r http.RoundTripper) http.RoundTripper {
	return hnynethttp.WrapRoundTripper(r)
}
