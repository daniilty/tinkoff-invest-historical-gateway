package service

import (
	schema "github.com/daniilty/tinkoff-invest-grpc-schema"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func normalizePBCandles(candles []*schema.Candle, interval uint32) []*schema.Candle {
	const minLen = 1

	if len(candles) < minLen {
		return candles
	}

	if interval == 0 {
		panic("0 interval")
	}

	for i := 1; i < len(candles); i++ {
		if !isSkippedToCandle(candles[i-1], candles[i]) {
			continue
		}

		counter := candles[i-1].GetTs().GetSeconds() + int64(interval)
		to := candles[i].GetTs().GetSeconds()

		for counter < to {
			candle := *candles[i-1]
			candle.Ts = timestamppb.New(convertTimestampToTime(uint32(counter)))

			candles = append(candles, &candle)

			counter += int64(interval)
		}
	}

	return candles
}

func isSkippedToCandle(current *schema.Candle, next *schema.Candle) bool {
	const skippedDays = 2 * 86400

	nextTS := next.GetTs().GetSeconds()
	currentTS := current.GetTs().GetSeconds()

	return nextTS-currentTS == skippedDays
}
