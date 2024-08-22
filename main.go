package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type assetsResponse struct {
	Data      []assetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}
type assetData struct {
	ID           string `json:"id"`
	Rank         string `json:"rank"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Supply       string `json:"supply"`
	MaxSupply    string `json:"maxSupply"`
	MarketCapUSD string `json:"marketCapUSD"`
	VolumeUSD24h string `json:"VolumeUSD24Hr"`
	PriceUSD     string `json:"priceUSD"`
}

func (d assetData) Info() string {
	return fmt.Sprintf("ID=%s Rank=%s Symbol=%s Name=%s Price=%s", d.ID, d.Rank, d.Symbol, d.Name, d.PriceUSD)
}

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "%s", time.Now().Format(time.DateTime))
	return l.next.RoundTrip(r)
}

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect")
			return nil
		},
		Transport: &loggingRoundTripper{
			logger: os.Stdout,
			next:   http.DefaultTransport,
		},
		Timeout: time.Second * 15,
	}

	resp, err := client.Get("https://api.coincap.io/v2/assets")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response :", resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var r assetsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		log.Fatal(err)
	}

	for _, asset := range r.Data {
		fmt.Println(asset.Info())
	}
}
