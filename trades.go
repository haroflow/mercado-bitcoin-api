package mercadobitcoin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Trade struct {
	// Transaction ID
	TID       int `json:"tid"`
	Timestamp int `json:"date"`
	Date      time.Time
	// Buy/Sell
	Type   string  `json:"type"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}

func (t *Trade) String() string {
	return fmt.Sprintf("TID: %d, Timestamp: %d, Type: %s, Price: %f, Amount: %f, Date: %s",
		t.TID, t.Timestamp, t.Type, t.Price, t.Amount, t.Date)
}

// GetTradesFilter is passed to GetTrades to filter by Timestamp or TID
type GetTradesFilter struct {
	FromTimestamp string
	ToTimestamp   string
	TID           string
}

// GetTrades gets the last 1000 trades for a coin. You can specify filters by Timestamp or TID using GetTradesFilter.
func GetTrades(coin Coin, filter *GetTradesFilter) ([]*Trade, error) {
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

	resp, err := http.Get(url)
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

	var trades []*Trade
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
