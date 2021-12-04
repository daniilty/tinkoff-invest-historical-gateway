package postgres

// Candle - stock candle data.
type Candle struct {
	ID           int
	FIGI         string
	TimeStamp    uint32
	HighestPrice float64
	LowestPrice  float64
	Open         float64
	Close        float64
}

// Candles - multiple candles.
type Candles []Candle
