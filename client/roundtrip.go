package client

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type LoggingRoundTripper struct {
	Logger io.Writer
	Next   http.RoundTripper
}

func (l LoggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.Logger, "[%s] %s %s\n", time.Now().Format(time.RFC3339), r.Method, r.URL)
	return l.Next.RoundTrip(r)
}
