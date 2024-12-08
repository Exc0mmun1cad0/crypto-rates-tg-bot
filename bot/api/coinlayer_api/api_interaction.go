package coinlayerapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	e "github.com/telegram_bot/bot/lib/error_wrapping"
)

func GetLatestRates() (res CoinlayerResponse, err error) {
	q := url.Values{}
	q.Add("access_key", os.Getenv("COINLAYER_API_KEY"))

	data, err := doRequest(q)
	if err != nil {
		return CoinlayerResponse{}, e.Wrap(errGettingRates, err)
	}

	var rates CoinlayerResponse
	err = json.Unmarshal(data, &rates)
	if err != nil {
		return CoinlayerResponse{}, e.Wrap(errGettingRates, err)
	}

	return rates, nil
}

func doRequest(query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   livePath,
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, e.Wrap(errDoingRequest, err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		return nil, e.Wrap(errDoingRequest, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, e.Wrap(errDoingRequest, err)
	}

	return body, nil
}