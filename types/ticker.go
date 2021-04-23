package types

import (
	"fmt"
	"time"
)

// Ticker is a summary of the last 24 hours
type Ticker struct {
	Coin        Coin
	Description string
	High        float64 `json:"high,string"`
	Low         float64 `json:"low,string"`
	Vol         float64 `json:"vol,string"`
	Last        float64 `json:"last,string"`
	Buy         float64 `json:"buy,string"`
	Sell        float64 `json:"sell,string"`
	Open        float64 `json:"open,string"`
	Timestamp   int     `json:"date"`
	Date        time.Time
}

func (t *Ticker) String() string {
	return fmt.Sprintf("Coin: %s, Description: %s, High: %f, Low: %f, Vol: %f, Last: %f, Buy: %f, Sell: %f, Open: %f, Timestamp: %d, Date: %v",
		t.Coin, t.Description, t.High, t.Low, t.Vol, t.Last, t.Buy, t.Sell, t.Open, t.Timestamp, t.Date)
}
