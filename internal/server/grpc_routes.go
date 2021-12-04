package server

import (
	"context"

	schema "github.com/daniilty/tinkoff-invest-grpc-schema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetCandles - get and download missing candles.
func (g *GRPC) GetCandles(ctx context.Context, req *schema.CandlesRequest) (*schema.CandlesResponse, error) {
	candles, err := g.downloader.GetCandles(ctx, convertGRPCRequestToService(req))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to download candles")
	}

	candlesGRPC := convertServiceCandlesToGRPC(candles)

	return &schema.CandlesResponse{
		Candles: candlesGRPC,
	}, nil
}
