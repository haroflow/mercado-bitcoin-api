package types

import (
	"fmt"
	"time"
)

type DaySummary struct {
	DateStr  string `json:"date"`
	Date     time.Time
	Opening  float64 `json:"opening"`
	Closing  float64 `json:"closing"`
	Lowest   float64 `json:"lowest"`
	Highest  float64 `json:"highest"`
	Volume   float64 `json:"volume"`
	Quantity float64 `json:"quantity"`
	Amount   float64 `json:"amount"`
	AvgPrice float64 `json:"avg_price"`
}

func (ds *DaySummary) String() string {
	return fmt.Sprintf("Date: %s, DateStr: %s, Opening: %f, Closing: %f, Lowest: %f, Highest: %f, Volume: %f, Quantity: %f, Amount: %f, AvgPrice: %f",
		ds.Date, ds.DateStr, ds.Opening, ds.Closing, ds.Lowest, ds.Highest, ds.Volume, ds.Quantity, ds.Amount, ds.AvgPrice,
	)
}
