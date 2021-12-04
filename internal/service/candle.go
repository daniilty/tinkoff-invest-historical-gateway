package service

type Candle struct {
	FIGI         string
	TimeStamp    uint32
	HighestPrice float64
	LowestPrice  float64
	Open         float64
	Close        float64
}

type Candles []Candle
