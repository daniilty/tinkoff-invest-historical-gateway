package service

import (
	"time"

	schema "github.com/daniilty/tinkoff-invest-grpc-schema"
	"github.com/daniilty/tinkoff-invest-historical-gateway/internal/data"
	"github.com/daniilty/tinkoff-invest-historical-gateway/internal/postgres"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertDBCandlesToInner(candles postgres.Candles) Candles {
	converted := make(Candles, 0, len(candles))

	for i := range candles {
		converted = append(converted, convertDBCandleToInner(candles[i]))
	}

	return converted
}

func convertDBCandleToInner(candle postgres.Candle) Candle {
	return Candle{
		FIGI:         candle.FIGI,
		TimeStamp:    candle.TimeStamp,
		LowestPrice:  candle.LowestPrice,
		HighestPrice: candle.HighestPrice,
		Open:         candle.Open,
		Close:        candle.Close,
	}
}

func convertPBCandleToInner(candle *schema.Candle, figi string) Candle {
	return Candle{
		FIGI:         figi,
		TimeStamp:    uint32(candle.GetTs().GetSeconds()),
		LowestPrice:  candle.GetOpenPrice(),
		HighestPrice: candle.GetHighPrice(),
		Open:         candle.GetOpenPrice(),
		Close:        candle.GetClosePrice(),
	}
}

func convertPBCandleToDB(candle *schema.Candle, figi string) postgres.Candle {
	return postgres.Candle{
		FIGI:         figi,
		TimeStamp:    uint32(candle.Ts.GetSeconds()),
		LowestPrice:  candle.GetLowPrice(),
		HighestPrice: candle.GetHighPrice(),
		Open:         candle.GetOpenPrice(),
		Close:        candle.GetClosePrice(),
	}
}

func convertDataIntervalsToCandlesRequests(intervals []*data.Interval, figi string) []*schema.CandlesRequest {
	converted := make([]*schema.CandlesRequest, 0, len(intervals))

	for i := range intervals {
		converted = append(converted, convertDataIntervalToCandlesRequest(intervals[i], figi))
	}

	return converted
}

func convertDataIntervalToCandlesRequest(interval *data.Interval, figi string) *schema.CandlesRequest {
	fromTime := convertTimestampToTime(interval.From)
	toTime := convertTimestampToTime(interval.To)

	return &schema.CandlesRequest{
		Figi:     figi,
		Interval: interval.Step,
		From:     timestamppb.New(fromTime),
		To:       timestamppb.New(toTime),
	}
}

func convertTimestampToTime(ts uint32) time.Time {
	const nsec = 0

	return time.Unix(int64(ts), nsec)
}
