package service

import (
	"fmt"
	"net/http"

	"github.com/haroflow/mercado-bitcoin-api/types"
)

// Default is used to request HTTP data from the API. Implements ServiceInterface.
type Default struct{}

func (d *Default) GetCoins() (*http.Response, error) {
	url := "https://www.mercadobitcoin.net/api/coins"
	return http.Get(url)
}

func (d *Default) GetTicker(coin types.Coin) (*http.Response, error) {
	url := "https://www.mercadobitcoin.net/api/" + string(coin) + "/ticker/"
	return http.Get(url)
}

func (d *Default) GetTrades(coin types.Coin, filter *GetTradesFilter) (*http.Response, error) {
	url := "https://www.mercadobitcoin.net/api/" + string(coin) + "/trades"
	if filter != nil {
		if filter.FromTimestamp != "" {
			url += fmt.Sprintf("/%s", filter.FromTimestamp)
		}
		if filter.FromTimestamp != "" && filter.ToTimestamp != "" {
			url += fmt.Sprintf("/%s", filter.ToTimestamp)
		}
		if filter.TID != "" {
			url += fmt.Sprintf("?tid=%s", filter.TID)
		}
	}

	return http.Get(url)
}

func (d *Default) GetDaySummary(coin types.Coin, day, month, year int) (*http.Response, error) {
	url := fmt.Sprintf("https://www.mercadobitcoin.net/api/%s/day-summary/%d/%d/%d", coin, year, month, day)
	return http.Get(url)
}

func (d *Default) GetOrderbook(coin types.Coin) (*http.Response, error) {
	url := fmt.Sprintf("https://www.mercadobitcoin.net/api/%s/orderbook", coin)
	return http.Get(url)
}
