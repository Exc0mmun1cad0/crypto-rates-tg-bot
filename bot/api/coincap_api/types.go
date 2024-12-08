package coincapapi

const (
	host = "api.coincap.io"
	path = "v2/assets"
)

const (
	errDoingRequest = "cannot do request"
	errGettingRates = "cannot get latest rates"
)

type Rates struct {
	Data      []CoinData `json:"data"`
	Timestamp int64      `json:"timestamp"`
}

type CoinData struct {
	ID     string `json:"id"`
	Rank   string `json:"rank"`
	Symbol string `json:"symbol"`
	//Name              string `json:"name"`
	//Supply            string `json:"supply"`
	//MaxSupply         string `json:"maxSupply"`
	//MarketCapUsd      string `json:"marketCapUsd"`
	//VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd string `json:"priceUsd"`
	//ChangePercent24Hr string `json:"changePercent24Hr"`
	//Vwap24Hr          string `json:"vwap24Hr"`
}
