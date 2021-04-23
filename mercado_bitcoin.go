// Data API for Mercado Bitcoin.
// https://www.mercadobitcoin.com.br/api-doc/
package mercadobitcoin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/haroflow/mercado-bitcoin-api/service"
	"github.com/haroflow/mercado-bitcoin-api/types"
)

// Client is used to access the API. Start with mercadobitcoin.NewClient().
type Client struct {
	Service service.ServiceInterface
}

// NewClient returns a new client with the default provider using HTTP.
func NewClient() *Client {
	return &Client{
		Service: &service.Default{},
	}
}

// GetTicker returns the 24-hour summary for the coin.
func (m *Client) GetTicker(coin types.Coin) (*types.Ticker, error) {
	resp, err := m.Service.GetTicker(coin)
	if err != nil {
		return nil, fmt.Errorf("error requesting ticker for coin %s: %s", coin, err)
	}
	defer resp.Body.Close()

	type tickerResponse struct {
		Ticker types.Ticker
	}
	var response tickerResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error decoding ticker for coin %s: %s: %s", coin, err, msg)
	}

	response.Ticker.Coin = coin
	response.Ticker.Description = types.Coins[coin]
	response.Ticker.Date = time.Unix(int64(response.Ticker.Timestamp), 0)

	return &response.Ticker, nil
}

// GetTrades gets the last 1000 trades for a coin. You can specify filters by Timestamp or TID using GetTradesFilter.
func (c *Client) GetTrades(coin types.Coin, filter *service.GetTradesFilter) ([]*types.Trade, error) {
	resp, err := c.Service.GetTrades(coin, filter)
	if err != nil {
		return nil, fmt.Errorf("error requesting trades for coin %s: %s", coin, err)
	}
	defer resp.Body.Close()

	// Example response:
	// [
	// 	{
	// 		"tid": 10007777, "date": 1618979594, "type": "sell", "price": 313100.00031, "amount": 0.00011856
	// 	},
	//  ...
	// ]

	var trades []*types.Trade
	err = json.NewDecoder(resp.Body).Decode(&trades)
	if err != nil {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error decoding trades for coin %s: %s: %s", coin, err, msg)
	}

	for i := range trades {
		trades[i].Date = time.Unix(int64(trades[i].Timestamp), 0)
	}

	return trades, nil
}
