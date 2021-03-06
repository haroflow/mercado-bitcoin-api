package mercadobitcoin_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	mercadobitcoin "github.com/haroflow/mercado-bitcoin-api"
	"github.com/haroflow/mercado-bitcoin-api/service"
	"github.com/haroflow/mercado-bitcoin-api/types"
)

var FakeResponseNotFound = func() (*http.Response, error) {
	return &http.Response{
		StatusCode: 404,
		Body:       ioutil.NopCloser(strings.NewReader(`Not Found`)),
	}, nil
}

var FakeResponse500 = func() (*http.Response, error) {
	return &http.Response{}, fmt.Errorf("Failed to GET response")
}

type StubMercadoBitcoinAPI struct {
	FakeGetCoins      func() (*http.Response, error)
	FakeGetTicker     func() (*http.Response, error)
	FakeGetTrades     func() (*http.Response, error)
	FakeGetDaySummary func() (*http.Response, error)
	FakeGetOrderbook  func() (*http.Response, error)
}

func (s *StubMercadoBitcoinAPI) GetCoins() (*http.Response, error) {
	return s.FakeGetCoins()
}

func (s *StubMercadoBitcoinAPI) GetTicker(coin types.Coin) (*http.Response, error) {
	return s.FakeGetTicker()
}

func (s *StubMercadoBitcoinAPI) GetTrades(coin types.Coin, filter *service.GetTradesFilter) (*http.Response, error) {
	return s.FakeGetTrades()
}

func (s *StubMercadoBitcoinAPI) GetDaySummary(coin types.Coin, day, month, year int) (*http.Response, error) {
	return s.FakeGetDaySummary()
}

func (s *StubMercadoBitcoinAPI) GetOrderbook(coin types.Coin) (*http.Response, error) {
	return s.FakeGetOrderbook()
}

func TestClientGetTicker(t *testing.T) {
	t.Run("return ticker for a valid coin name", func(t *testing.T) {
		fakeResponse := func() (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader(`{"ticker": {"high":"318299.99997000","low":"299999.99990000","vol":"171.17893075","last":"302006.99104000","buy":"302006.99104000","sell":"302007.00000000","open":"317651.00000000","date":1619058701}}`)),
			}, nil
		}

		api := &mercadobitcoin.Client{
			Service: &StubMercadoBitcoinAPI{
				FakeGetTicker: fakeResponse,
			},
		}

		resp, err := api.GetTicker("BTC")

		assertNoError(t, err)
		if resp == nil {
			t.Fatal("expected response, got nil")
		}
		if resp.Last == 0 {
			t.Error("last price should not be zero")
		}
		if resp.Coin == "" {
			t.Errorf("expected the name of the coin, got empty string")
		}
		if resp.Description == "" {
			t.Errorf("expected the description of the coin, got empty string")
		}
		if resp.Date.IsZero() {
			t.Errorf("expected the date of the response, got empty time.Time")
		}
	})

	t.Run("return error 'Not Found' for an invalid coin name", func(t *testing.T) {
		api := &mercadobitcoin.Client{
			Service: &StubMercadoBitcoinAPI{
				FakeGetTicker: FakeResponseNotFound,
			},
		}

		resp, err := api.GetTicker("123BTC")

		assertError(t, err)
		if resp != nil {
			t.Fatalf("didnt expected response, got %v", resp)
		}
	})

	t.Run("return error on http failure", func(t *testing.T) {
		api := &mercadobitcoin.Client{
			Service: &StubMercadoBitcoinAPI{
				FakeGetTicker: FakeResponse500,
			},
		}

		resp, err := api.GetTicker("")
		if resp != nil {
			t.Errorf("didnt expected a response, got %v", resp)
		}

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func TestClientGetTrades(t *testing.T) {
	t.Run("should return trades for a valid coin", func(t *testing.T) {
		fakeResponse := func() (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader(`[{"tid":10062662,"date":1619134035,"type":"buy","price":284999.99,"amount":0.01293581},{"tid":10062663,"date":1619134037,"type":"sell","price":284999.8901,"amount":0.0103513},{"tid":10062664,"date":1619134040,"type":"buy","price":284999.99,"amount":0.00175438},{"tid":10062665,"date":1619134040,"type":"buy","price":284999.99,"amount":0.00350877},{"tid":10062666,"date":1619134042,"type":"buy","price":284999.99,"amount":0.00392287},{"tid":10062667,"date":1619134043,"type":"buy","price":284999.99,"amount":0.00687512},{"tid":10062668,"date":1619134045,"type":"buy","price":284999.99,"amount":0.00052631},{"tid":10062669,"date":1619134045,"type":"buy","price":284999.99,"amount":0.00035087},{"tid":10062670,"date":1619134047,"type":"buy","price":284999.99,"amount":0.00070175},{"tid":10062671,"date":1619134053,"type":"buy","price":284999.99,"amount":0.00017543},{"tid":10062672,"date":1619134054,"type":"buy","price":284999.99,"amount":0.01052631},{"tid":10062673,"date":1619134055,"type":"buy","price":284999.99,"amount":0.00180822},{"tid":10062674,"date":1619134055,"type":"buy","price":285000,"amount":0.0001},{"tid":10062675,"date":1619134055,"type":"buy","price":285000,"amount":0.00245009},{"tid":10062676,"date":1619134056,"type":"buy","price":284999.75,"amount":0.00008737},{"tid":10062677,"date":1619134059,"type":"buy","price":284999.75,"amount":0.00008525},{"tid":10062678,"date":1619134059,"type":"buy","price":285000,"amount":0.00166913},{"tid":10062679,"date":1619134060,"type":"buy","price":285000,"amount":0.00007017}]`)),
			}, nil
		}

		api := &mercadobitcoin.Client{
			Service: &StubMercadoBitcoinAPI{
				FakeGetTrades: fakeResponse,
			},
		}

		trades, err := api.GetTrades("BTC", nil)
		if err != nil {
			t.Fatalf("didnt expected an error, got %s", err)
		}
		if len(trades) == 0 {
			t.Errorf("expected response, got %v", trades)
		}
	})

	t.Run("return error 'Not Found' for an invalid coin name", func(t *testing.T) {
		api := &mercadobitcoin.Client{
			Service: &StubMercadoBitcoinAPI{
				FakeGetTrades: FakeResponseNotFound,
			},
		}

		resp, err := api.GetTrades("123BTC", nil)

		assertError(t, err)
		if resp != nil {
			t.Fatalf("didnt expected response, got %v", resp)
		}
	})

	t.Run("return error on http failure", func(t *testing.T) {
		api := &mercadobitcoin.Client{
			Service: &StubMercadoBitcoinAPI{
				FakeGetTrades: FakeResponse500,
			},
		}

		resp, err := api.GetTrades("BTC", nil)

		assertError(t, err)
		if resp != nil {
			t.Fatalf("didnt expected response, got %v", resp)
		}
	})
}

func TestClientGetDaySummary(t *testing.T) {
	t.Run("should return day summary when called with a valid coin", func(t *testing.T) {
		fakeResponse := func() (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader(`{"date":"2020-02-01","opening":40009.09990999,"closing":39755,"lowest":39700,"highest":40139.98,"volume":1557826.34691214,"quantity":39.09361166,"amount":1668,"avg_price":39848.61671161}`)),
			}, nil
		}

		api := &mercadobitcoin.Client{
			Service: &StubMercadoBitcoinAPI{
				FakeGetDaySummary: fakeResponse,
			},
		}

		day, month, year := 1, 2, 2020
		resp, err := api.GetDaySummary("BTC", day, month, year)

		if err != nil {
			t.Fatalf("didnt expected an error, got %s", err)
		}
		if resp == nil {
			t.Fatal("expected response, got nil")
		}

		if resp.Date.IsZero() {
			t.Errorf("expected date, got %s", resp.Date)
		}
	})

	t.Run("should return an error when called with an invalid coin", func(t *testing.T) {
		api := &mercadobitcoin.Client{
			Service: &StubMercadoBitcoinAPI{
				FakeGetDaySummary: FakeResponseNotFound,
			},
		}

		day, month, year := 1, 2, 2020
		resp, err := api.GetDaySummary("123BTC123", day, month, year)

		assertError(t, err)
		if resp != nil {
			t.Errorf("didnt expect a response, got %v", resp)
		}
	})

	t.Run("return error on http failure", func(t *testing.T) {
		api := &mercadobitcoin.Client{
			Service: &StubMercadoBitcoinAPI{
				FakeGetDaySummary: FakeResponse500,
			},
		}

		day, month, year := 1, 2, 2020
		resp, err := api.GetDaySummary("BTC", day, month, year)

		assertError(t, err)
		if resp != nil {
			t.Fatalf("didnt expected response, got %v", resp)
		}
	})
}

func TestClientGetOrderbook(t *testing.T) {
	orderbookTestData, err := ioutil.ReadFile("./testdata/orderbook.json")
	if err != nil {
		t.Fatal("could not read testdata/orderbook.json")
	}

	fakeResponse := func() (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(strings.NewReader(string(orderbookTestData))),
		}, nil
	}

	api := &mercadobitcoin.Client{
		Service: &StubMercadoBitcoinAPI{
			FakeGetOrderbook: fakeResponse,
		},
	}

	res, err := api.GetOrderbook("BTC")
	if err != nil {
		t.Fatalf("didn't expected error, got %q", err)
	}
	if res == nil {
		t.Fatal("expected response, got nil")
	}
	if len(res.Asks) == 0 {
		t.Errorf("expected at least one Ask, got none")
	}
	if len(res.Bids) == 0 {
		t.Errorf("expected at least one Bid, got none")
	}
}

func TestClientGetCoins(t *testing.T) {
	fakeResponse := func() (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(strings.NewReader(`["BCH","BTC","ETH","LTC","XRP","MBPRK01","MBPRK02","MBPRK03","MBPRK04","MBCONS01","USDC","WBX","CHZ","MBCONS02","PAXG","MBVASCO01","LINK","PSGFT","JUVFT","ASRFT","ATMFT","GALFT","CAIFT","MCO2","ACMFT","OGFT"]`)),
		}, nil
	}

	api := &mercadobitcoin.Client{
		Service: &StubMercadoBitcoinAPI{
			FakeGetCoins: fakeResponse,
		},
	}

	res, err := api.GetCoins()
	if err != nil {
		t.Fatalf("didn't expected error, got %q", err)
	}
	if res == nil {
		t.Fatal("expected response, got nil")
	}
	if len(res) == 0 {
		t.Errorf("expected at least one Coin, got none")
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("expected an error, got %q", err)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("didnt expected an error, got %q", err)
	}
}
