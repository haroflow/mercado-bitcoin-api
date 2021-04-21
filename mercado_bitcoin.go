// Data API for Mercado Bitcoin
// https://www.mercadobitcoin.com.br/api-doc/
package mercadobitcoin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GetTicker solicita o resumo das últimas 24 horas de negociação da moeda.
func GetTicker(coin Coin) (*Ticker, error) {
	url := "https://www.mercadobitcoin.net/api/" + string(coin) + "/ticker/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error requesting ticker: %s", err)
	}
	defer resp.Body.Close()

	// The response we want is inside the "ticker" key, so we create an anonymous struct to get just that:
	// {
	//   "ticker": {
	//     "high": "323000.00000000", "low": "307012.36000000",
	//     "vol": "166.93637836", "last": "312000.00317000",
	//     "buy": "312000.00000000", "sell": "312000.00315000",
	//     "open": "313862.84825000", "date": 1618992914
	//   }
	// }
	type tickerResponse struct {
		Ticker Ticker
	}

	var response tickerResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error decoding ticker: %s", msg)
	}

	response.Ticker.Coin = coin
	response.Ticker.Description = Coins[coin]
	response.Ticker.Date = time.Unix(int64(response.Ticker.Timestamp), 0)

	return &response.Ticker, nil
}
