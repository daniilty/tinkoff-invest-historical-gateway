package service

import (
	"context"

	schema "github.com/daniilty/tinkoff-invest-grpc-schema"
	"github.com/daniilty/tinkoff-invest-historical-gateway/internal/postgres"
)

type Downloader interface {
	GetCandles(context.Context, CandlesRequest) (Candles, error)
}

type DownloaderImpl struct {
	db     postgres.DB
	client schema.TinkoffInvestAPIGatewayClient
}

func NewDownloaderImpl(db postgres.DB, client schema.TinkoffInvestAPIGatewayClient) *DownloaderImpl {
	return &DownloaderImpl{
		db:     db,
		client: client,
	}
}
