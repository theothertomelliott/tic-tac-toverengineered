package otelmonitoring

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// WrapHTTP will add monitoring middleware to an http handler
func (m *Monitoring) WrapHTTP(handler http.Handler) http.Handler {
	return otelhttp.NewHandler(handler, "monitor")
}

// WrapHTTPTransport will add monitoring middleware to an HTTP Transport.
// This allows for trace propagation.
func (m *Monitoring) WrapHTTPTransport(r http.RoundTripper) http.RoundTripper {
	return otelhttp.NewTransport(r)
}
