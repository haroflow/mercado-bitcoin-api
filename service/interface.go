package service

import (
	"net/http"

	"github.com/haroflow/mercado-bitcoin-api/types"
)

// ServiceInterface is used to implement a Mercado Bitcoin API service.
type ServiceInterface interface {
	GetCoins() (*http.Response, error)
	GetTicker(coin types.Coin) (*http.Response, error)
	GetTrades(coin types.Coin, filter *GetTradesFilter) (*http.Response, error)
	GetDaySummary(coin types.Coin, day, month, year int) (*http.Response, error)
	GetOrderbook(coin types.Coin) (*http.Response, error)
}

// GetTradesFilter is passed to GetTrades to filter by Timestamp or TID
type GetTradesFilter struct {
	FromTimestamp string
	ToTimestamp   string
	TID           string
}
