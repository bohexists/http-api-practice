package coincap

import "fmt"

// Responses for CoinCap API
type assetsResponse struct {
	Assets    []Asset `json:"data"`
	Timestamp int64   `json:"timestamp"`
}

// Responses for CoinCap API
type assetResponse struct {
	Asset     Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

// Asset represents the detailed information about an asset returned by the CoinCap API.
type Asset struct {
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

// Info returns the detailed information about the asset as a string.
func (d Asset) Info() string {
	return fmt.Sprintf("ID=%s Rank=%s Symbol=%s Name=%s Price=%s", d.ID, d.Rank, d.Symbol, d.Name, d.PriceUSD)
}
