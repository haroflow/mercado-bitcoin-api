package main

import (
	"fmt"
	"os"

	mercadobitcoin "github.com/haroflow/mercado-bitcoin-api"
)

func main() {
	fmt.Println("Mercado Bitcoin API")
	fmt.Println()

	fmt.Println("# Ticker example:")
	coins := []mercadobitcoin.Coin{
		"BTC",
		"LTC",
		"CHZ",
		"XRP",
	}

	for _, coin := range coins {
		ticker, err := mercadobitcoin.GetTicker(coin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%-10s R$ %14f\n", ticker.Coin, ticker.Last)
	}
	fmt.Println()

	fmt.Println("# Trades example:")
	trades, err := mercadobitcoin.GetTrades("BTC", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, trade := range trades[:5] {
		fmt.Println(trade)
	}
	fmt.Printf("Total trades returned: %d\n\n", len(trades))

	fmt.Println("# Trades with timestamp after 1501871369:")
	trades, err = mercadobitcoin.GetTrades("BTC", &mercadobitcoin.GetTradesFilter{
		FromTimestamp: "1501871369",
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, trade := range trades[:5] {
		fmt.Println(trade)
	}
	fmt.Printf("Total trades returned: %d\n\n", len(trades))

	fmt.Println("# Trades with timestamp between 1501871369 and 1501871388:")
	trades, err = mercadobitcoin.GetTrades("BTC", &mercadobitcoin.GetTradesFilter{
		FromTimestamp: "1501871369",
		ToTimestamp:   "1501871388",
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, trade := range trades {
		fmt.Println(trade)
	}
	fmt.Printf("Total trades returned: %d\n\n", len(trades))

	fmt.Println("# Trades with Transaction ID after 5000:")
	trades, err = mercadobitcoin.GetTrades("BTC", &mercadobitcoin.GetTradesFilter{
		TID: "5000",
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, trade := range trades[:5] {
		fmt.Println(trade)
	}
	fmt.Printf("Total trades returned: %d\n\n", len(trades))

}
