package ipstack

type IPInfo struct {
	IP            string     `json:"ip"`
	Type          string     `json:"type"`
	ContinentCode string     `json:"continent_code"`
	ContinentName string     `json:"continent_name"`
	CountryCode   string     `json:"country_code"`
	CountryName   string     `json:"country_name"`
	RegionCode    string     `json:"region_code"`
	RegionName    string     `json:"region_name"`
	City          string     `json:"city"`
	Zip           string     `json:"zip"`
	Latitude      float64    `json:"latitude"`
	Longitude     float64    `json:"longitude"`
	TimeZone      TimeZone   `json:"time_zone"`
	Currency      Currency   `json:"currency"`
	Connection    Connection `json:"connection"`
}

type TimeZone struct {
	ID               string `json:"id"`
	CurrentTime      string `json:"current_time"`
	GMTOffset        int    `json:"gmt_offset"`
	Code             string `json:"code"`
	IsDaylightSaving bool   `json:"is_daylight_saving"`
}

type Currency struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	SymbolNative string `json:"symbol_native"`
}

type Connection struct {
	ASN int    `json:"asn"`
	ISP string `json:"isp"`
}
