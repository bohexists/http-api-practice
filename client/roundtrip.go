package client

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// LoggingRoundTripper is a custom implementation of http.RoundTripper
// that logs each HTTP request before passing it to the next RoundTripper.
type LoggingRoundTripper struct {
	Logger io.Writer         // Logger is an io.Writer where log messages will be written.
	Next   http.RoundTripper // Next is the next RoundTripper in the chain to which the request will be forwarded.
}

// RoundTrip executes a single HTTP transaction, logging the request details
// and then passing the request to the next RoundTripper in the chain.
func (l LoggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	// Log the HTTP request with the current timestamp, method, and URL.
	fmt.Fprintf(l.Logger, "[%s] %s %s\n", time.Now().Format(time.RFC3339), r.Method, r.URL)
	// Forward the request to the next RoundTripper and return its response.
	return l.Next.RoundTrip(r)
}
