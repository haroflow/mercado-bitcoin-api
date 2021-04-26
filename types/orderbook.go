package types

type Orderbook struct {
	Asks      [][]float64 // TODO change these into structs with Price and Quantity
	Bids      [][]float64
	Timestamp int64
}
