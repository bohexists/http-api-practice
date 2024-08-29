package coincap

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bohexists/http-api-practice/client/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Client is a structure that holds an HTTP client for making requests to the CoinCap API.
type Client struct {
	client *http.Client
}

// NewClient creates a new CoinCap API client with a specified timeout.
// If the timeout is 0, an error is returned.
func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout 0")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &middleware.LoggingRoundTripper{
				Logger: os.Stdout,
				Next:   http.DefaultTransport,
			},
		},
	}, nil
}

// GetAssets retrieves a list of all assets from the CoinCap API.
func (c Client) GetAssets() ([]Asset, error) {
	resp, err := c.client.Get("https://api.coincap.io/v2/assets")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var r assetsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r.Assets, nil
}

// GetAsset retrieves information about a specific asset from the CoinCap API.
func (c Client) GetAsset(name string) (Asset, error) {
	url := fmt.Sprintf("https://api.coincap.io/v2/assets/%s", name)
	resp, err := c.client.Get(url)
	if err != nil {
		return Asset{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Asset{}, err
	}

	var r assetResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return Asset{}, err
	}

	return r.Asset, nil
}
