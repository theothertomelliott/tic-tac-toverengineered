package logmonitoring

import (
	"log"
	"net/http"
)

// WrapHTTP will add monitoring middleware to an http handler
func (m *Monitoring) WrapHTTP(handler http.Handler) http.Handler {
	wrappedHandler := func(w http.ResponseWriter, r *http.Request) {
		log.Println("http request: ", r.URL.Path)
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(wrappedHandler)
}

// WrapHTTPTransport will add monitoring middleware to an HTTP Transport.
// This allows for trace propagation.
func (m *Monitoring) WrapHTTPTransport(r http.RoundTripper) http.RoundTripper {
	return &logRT{rt: r}
}

var _ http.RoundTripper = &logRT{}

// logRT wraps an http.RoundTripper with logging
type logRT struct {
	rt http.RoundTripper
}

func (l *logRT) RoundTrip(r *http.Request) (*http.Response, error) {
	log.Println("http request: ", r.URL.Path)
	return l.rt.RoundTrip(r)
}
