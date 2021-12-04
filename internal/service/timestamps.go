package service

func getCandlesTimestamps(candles Candles) []uint32 {
	tt := make([]uint32, 0, len(candles))

	for i := range candles {
		tt = append(tt, candles[i].TimeStamp)
	}

	return tt
}
