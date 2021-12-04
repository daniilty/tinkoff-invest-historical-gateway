package service

import schema "github.com/daniilty/tinkoff-invest-grpc-schema"

func copyGRPCCandle(candle *schema.Candle) *schema.Candle {
	return &schema.Candle{
		LowPrice:   candle.GetLowPrice(),
		HighPrice:  candle.GetHighPrice(),
		OpenPrice:  candle.GetOpenPrice(),
		ClosePrice: candle.GetClosePrice(),
		Volume:     candle.GetVolume(),
		Ts:         candle.GetTs(),
	}
}
