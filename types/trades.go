package types

import (
	"fmt"
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
