package domain

type USD float32
type Time int64 // Unix Milliseconds
type StockSymbol string

type Candle struct {
	open      USD
	close     USD
	high      USD
	low       USD
	timestamp Time
	volume    uint64
	symbol    StockSymbol
}
