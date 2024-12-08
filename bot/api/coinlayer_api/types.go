package coinlayerapi

const (
	host     = `api.coinlayer.com`
	livePath = `api/live`
)

const (
	errDoingRequest = "cannot do request"
	errGettingRates = "cannot get latest rates"
)

type CoinlayerResponse struct {
	Success   bool   `json:"success"`
	Timestamp int64  `json:"timestamp"`
	Target    string `json:"target"`
	Rates     Rates  `json:"rates"`
}

type Rates struct {
	Bitcoin   float64 `json:"BTC"`
	Ethereum  float64 `json:"ETH"`
	Monero    float64 `json:"XMR"`
	Dogecoin  float64 `json:"DOGE"`
	Chainlink float64 `json:"LINK"`
}
