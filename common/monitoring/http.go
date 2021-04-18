package monitoring

import (
	"net/http"
)

// WrapHTTP will add monitoring middleware to an http handler
func WrapHTTP(handler http.Handler, name string) http.Handler {
	return Default.WrapHTTP(handler, name)
}

// WrapHTTPTransport will add monitoring middleware to an HTTP Transport.
// This allows for trace propagation.
func WrapHTTPTransport(r http.RoundTripper) http.RoundTripper {
	return Default.WrapHTTPTransport(r)
}
