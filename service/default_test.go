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

func TestGetDaySummary(t *testing.T) {
	t.Run("should get status 200 for a valid coin", func(t *testing.T) {
		s := &service.Default{}
		day, month, year := 1, 2, 2020
		resp, err := s.GetDaySummary("BTC", day, month, year)

		assertNoError(t, err)
		if resp == nil {
			t.Fatal("expected a response, got nil")
		}
		if resp.StatusCode != 200 {
			t.Errorf("expected status 200, got %d", resp.StatusCode)
		}
	})

	t.Run("should get status 404 for an invalid coin", func(t *testing.T) {
		s := &service.Default{}
		day, month, year := 1, 2, 2020
		resp, err := s.GetDaySummary("123BTC321", day, month, year)

		assertNoError(t, err)
		if resp == nil {
			t.Fatal("expected a response, got nil")
		}
		if resp.StatusCode != 404 {
			t.Errorf("expected status 404, got %d", resp.StatusCode)
		}
	})

	t.Run("should get status 500 for an invalid date", func(t *testing.T) {
		s := &service.Default{}
		day, month, year := 0, 0, 0
		resp, err := s.GetDaySummary("BTC", day, month, year)

		assertNoError(t, err)
		if resp == nil {
			t.Fatal("expected a response, got nil")
		}
		if resp.StatusCode != 500 {
			t.Errorf("expected status 500, got %d", resp.StatusCode)
		}
	})
}

func TestGetOrderbook(t *testing.T) {
	t.Run("should get status 200 for a valid coin", func(t *testing.T) {
		s := &service.Default{}
		resp, err := s.GetOrderbook("BTC")

		assertNoError(t, err)
		if resp == nil {
			t.Fatal("expected a response, got nil")
		}
		if resp.StatusCode != 200 {
			t.Errorf("expected status 200, got %d", resp.StatusCode)
		}
	})

	t.Run("should get status 404 for an invalid coin", func(t *testing.T) {
		s := &service.Default{}
		resp, err := s.GetOrderbook("123BTC321")

		assertNoError(t, err)
		if resp == nil {
			t.Fatal("expected a response, got nil")
		}
		if resp.StatusCode != 404 {
			t.Errorf("expected status 404, got %d", resp.StatusCode)
		}
	})
}

func TestGetCoins(t *testing.T) {
	s := &service.Default{}
	resp, err := s.GetCoins()

	assertNoError(t, err)
	if resp == nil {
		t.Fatal("expected a response, got nil")
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("didnt expected an error, got %q", err)
	}
}
