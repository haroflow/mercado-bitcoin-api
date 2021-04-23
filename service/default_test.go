package service_test

import (
	"testing"

	"github.com/haroflow/mercado-bitcoin-api/service"
)

func TestGetTicker(t *testing.T) {
	t.Run("existing coin should respond OK", func(t *testing.T) {
		s := &service.Default{}
		got, err := s.GetTicker("BTC")

		assertNoError(t, err)

		if got == nil {
			t.Fatal("expected a response, got nil")
		}
		if got.StatusCode != 200 {
			t.Errorf("expected 200, got %d", got.StatusCode)
		}
	})

	t.Run("non existing coin should error", func(t *testing.T) {
		s := &service.Default{}
		got, err := s.GetTicker("AAAAAAAAAAAA")

		assertNoError(t, err)

		if got == nil {
			t.Fatal("expected a response, got nil")
		}
		if got.StatusCode != 404 {
			t.Errorf("expected 404, got %d", got.StatusCode)
		}

	})
}

func TestGetTrades(t *testing.T) {
	t.Run("existing coin should respond OK", func(t *testing.T) {
		s := &service.Default{}
		got, err := s.GetTrades("BTC", nil)

		assertNoError(t, err)

		if got == nil {
			t.Fatal("expected a response, got nil")
		}
		if got.StatusCode != 200 {
			t.Errorf("expected 200, got %d", got.StatusCode)
		}
	})

	t.Run("non existing coin should error", func(t *testing.T) {
		s := &service.Default{}
		got, err := s.GetTrades("AAAAAAAAAAAA", nil)

		assertNoError(t, err)
		if got == nil {
			t.Fatal("expected a response, got nil")
		}
		if got.StatusCode != 404 {
			t.Errorf("expected 404, got %d", got.StatusCode)
		}
	})

	t.Run("trades after timestamp", func(t *testing.T) {
		s := &service.Default{}
		got, err := s.GetTrades("BTC", &service.GetTradesFilter{
			FromTimestamp: "1501871369",
		})

		assertNoError(t, err)

		if got == nil {
			t.Fatal("expected a response, got nil")
		}
		if got.StatusCode != 200 {
			t.Errorf("expected 200, got %d", got.StatusCode)
		}
	})

	t.Run("trades between two timestamps", func(t *testing.T) {
		s := &service.Default{}
		got, err := s.GetTrades("BTC", &service.GetTradesFilter{
			FromTimestamp: "1501871369",
			ToTimestamp:   "1501891200",
		})

		assertNoError(t, err)

		if got == nil {
			t.Fatal("expected a response, got nil")
		}
		if got.StatusCode != 200 {
			t.Errorf("expected 200, got %d", got.StatusCode)
		}
	})

	t.Run("trades after TID 5000", func(t *testing.T) {
		s := &service.Default{}
		got, err := s.GetTrades("BTC", &service.GetTradesFilter{
			TID: "5000",
		})

		assertNoError(t, err)

		if got == nil {
			t.Fatal("expected a response, got nil")
		}
		if got.StatusCode != 200 {
			t.Errorf("expected 200, got %d", got.StatusCode)
		}
	})
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("didnt expected an error, got %q", err)
	}
}
