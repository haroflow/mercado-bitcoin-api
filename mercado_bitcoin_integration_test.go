package mercadobitcoin_test

import (
	"testing"

	mercadobitcoin "github.com/haroflow/mercado-bitcoin-api"
	"github.com/haroflow/mercado-bitcoin-api/service"
)

func TestGetTicker(t *testing.T) {
	t.Run("existing coin should respond OK", func(t *testing.T) {
		api := mercadobitcoin.NewClient()
		got, err := api.GetTicker("BTC")

		assertNoError(t, err)

		if got == nil {
			t.Errorf("expected a response, got nil")
		}
	})

	t.Run("non existing coin should error", func(t *testing.T) {
		api := mercadobitcoin.NewClient()
		got, err := api.GetTicker("AAAAAAAAAAAA")

		assertError(t, err)

		if got != nil {
			t.Errorf("didnt expected a response, got %v", got)
		}
	})
}

func TestGetTrades(t *testing.T) {
	t.Run("existing coin should respond OK", func(t *testing.T) {
		api := mercadobitcoin.NewClient()
		got, err := api.GetTrades("BTC", nil)

		assertNoError(t, err)

		if got == nil {
			t.Errorf("expected a response, got nil")
		}

		if len(got) == 0 {
			t.Errorf("expected to receive at least one trade, got %v", got)
		}
	})

	t.Run("non existing coin should error", func(t *testing.T) {
		api := mercadobitcoin.NewClient()
		got, err := api.GetTrades("AAAAAAAAAAAA", nil)

		assertError(t, err)

		if got != nil || len(got) > 0 {
			t.Errorf("didnt expected a response, got %v", got)
		}
	})

	t.Run("trades after timestamp", func(t *testing.T) {
		api := mercadobitcoin.NewClient()
		got, err := api.GetTrades("BTC", &service.GetTradesFilter{
			FromTimestamp: "1501871369",
		})

		assertNoError(t, err)

		if got == nil {
			t.Errorf("expected a response, got nil")
		}

		if len(got) == 0 {
			t.Errorf("expected to receive at least one trade, got %v", got)
		}
	})

	t.Run("trades between two timestamps", func(t *testing.T) {
		api := mercadobitcoin.NewClient()
		got, err := api.GetTrades("BTC", &service.GetTradesFilter{
			FromTimestamp: "1501871369",
			ToTimestamp:   "1501891200",
		})

		assertNoError(t, err)

		if got == nil {
			t.Errorf("expected a response, got nil")
		}

		if len(got) == 0 {
			t.Errorf("expected to receive at least one trade, got %v", got)
		}
	})

	t.Run("trades after TID 5000", func(t *testing.T) {
		api := mercadobitcoin.NewClient()
		got, err := api.GetTrades("BTC", &service.GetTradesFilter{
			TID: "5000",
		})

		assertNoError(t, err)

		if got == nil {
			t.Errorf("expected a response, got nil")
		}

		if len(got) == 0 {
			t.Errorf("expected to receive at least one trade, got %v", got)
		}
	})
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("didnt expected an error, got %q", err)
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("expected an error, got %q", err)
	}
}
