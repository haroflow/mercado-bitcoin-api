package mercadobitcoin_test

import (
	"testing"

	mercadobitcoin "github.com/haroflow/mercado-bitcoin-api"
)

func TestGetTicker(t *testing.T) {
	t.Run("existing coin should respond OK", func(t *testing.T) {
		got, err := mercadobitcoin.GetTicker("BTC")

		assertNoError(t, err)

		if got == nil {
			t.Errorf("expected a response, got nil")
		}
	})

	t.Run("non existing coin should error", func(t *testing.T) {
		got, err := mercadobitcoin.GetTicker("AAAAAAAAAAAA")

		assertError(t, err)

		if got != nil {
			t.Errorf("didnt expected a response, got %v", got)
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

func assertGotResponse(t testing.TB, response string) {
	t.Helper()
	if response == "" {
		t.Errorf("should have received a response, got %q", response)
	}
}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNotNil(t testing.TB, got interface{}) {
	t.Helper()
	if got == nil {
		t.Errorf("got nil")
	}
}
