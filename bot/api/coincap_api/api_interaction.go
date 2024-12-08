package coincapapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	e "github.com/telegram_bot/bot/lib/error_wrapping"
)

func GetLatestRates(tokenList []string) (Rates, error) {
	q := url.Values{}
	q.Add("ids", strings.Join(tokenList, ","))

	data, err := doRequest(q)
	if err != nil {
		return Rates{}, e.Wrap(errGettingRates, err)
	}

	var rates Rates
	err = json.Unmarshal(data, &rates)
	if err != nil {
		return Rates{}, e.Wrap(errGettingRates, err)
	}

	return rates, nil
}

func doRequest(query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
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

