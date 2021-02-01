package honeycomb

import (
	"net/http"

	"github.com/honeycombio/beeline-go/wrappers/hnynethttp"
)

// WrapHTTP will add monitoring middleware to an http handler
func (m *Monitoring) WrapHTTP(handler http.Handler) http.Handler {
	return hnynethttp.WrapHandler(handler)
}

// WrapHTTPTransport will add monitoring middleware to an HTTP Transport.
// This allows for trace propagation.
func (m *Monitoring) WrapHTTPTransport(r http.RoundTripper) http.RoundTripper {
	return hnynethttp.WrapRoundTripper(r)
}
