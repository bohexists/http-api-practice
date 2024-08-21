package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l loggingRoundTripper) RoundTripper(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "%s", time.Now().Format(time.DateTime))
	return l.next.RoundTrip(r)
}

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect")
			return nil
		},
	}

	resp, err := client.Get("http://github.com/bohexists")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response :", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
