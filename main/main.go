package main

import (
	"fmt"

	mercadobitcoin "github.com/haroflow/mercado-bitcoin-api"
)

func main() {
	fmt.Println("Mercado Bitcoin API")

	coins := []mercadobitcoin.Coin{
		"BTC",
		"LTC",
		"CHZ",
		"XRP",
	}

	for _, coin := range coins {
		r, err := mercadobitcoin.GetTicker(coin)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%-10s R$ %14f\n", r.Coin, r.Last)
	}
}
