package ipstack

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bohexists/http-api-practice/client"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

func NewClient(apiKey string, timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout must be greater than 0")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &client.LoggingRoundTripper{
				Logger: os.Stdout,
				Next:   http.DefaultTransport,
			},
		},
		apiKey:  apiKey,
		baseURL: "http://api.ipstack.com/",
	}, nil
}

func (c *Client) GetIPInfo(ip string) (IPInfo, error) {
	url := fmt.Sprintf("%s%s?access_key=%s", c.baseURL, ip, c.apiKey)
	resp, err := c.client.Get(url)
	if err != nil {
		return IPInfo{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return IPInfo{}, err
	}

	var ipInfo IPInfo
	if err = json.Unmarshal(body, &ipInfo); err != nil {
		return IPInfo{}, err
	}

	return ipInfo, nil
}
