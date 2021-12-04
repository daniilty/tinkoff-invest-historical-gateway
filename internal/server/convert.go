package server

import (
	"time"

	schema "github.com/daniilty/tinkoff-invest-grpc-schema"
	"github.com/daniilty/tinkoff-invest-historical-gateway/internal/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertGRPCRequestToService(req *schema.CandlesRequest) service.CandlesRequest {
	return service.CandlesRequest{
		FIGI:     req.GetFigi(),
		From:     convertTimestamppbToService(req.GetFrom()),
		To:       convertTimestamppbToService(req.GetTo()),
		Interval: req.GetInterval(),
	}
}

func convertTimestamppbToService(ts *timestamppb.Timestamp) uint32 {
	return uint32(ts.GetSeconds())
}

func convertTimestampToTimestamppb(ts uint32) *timestamppb.Timestamp {
	tsTime := time.Unix(int64(ts), 0)

	return timestamppb.New(tsTime)
}

func convertServiceCandlesToGRPC(candles service.Candles) []*schema.Candle {
	converted := make([]*schema.Candle, 0, len(candles))

	for i := range candles {
		converted = append(converted, convertServiceCandleToGRPC(candles[i]))
	}

	return converted
}

func convertServiceCandleToGRPC(candle service.Candle) *schema.Candle {
	return &schema.Candle{
		Ts:         convertTimestampToTimestamppb(candle.TimeStamp),
		ClosePrice: candle.Close,
		OpenPrice:  candle.Open,
		HighPrice:  candle.HighestPrice,
		LowPrice:   candle.LowestPrice,
	}
}
