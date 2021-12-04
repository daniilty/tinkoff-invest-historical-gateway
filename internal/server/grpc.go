package server

import (
	schema "github.com/daniilty/tinkoff-invest-grpc-schema"
	"github.com/daniilty/tinkoff-invest-historical-gateway/internal/service"
	"go.uber.org/zap"
)

// GRPC - grpc server.
type GRPC struct {
	schema.UnimplementedTinkoffInvestHistoricalGatewayServer

	downloader service.Downloader
	logger     *zap.SugaredLogger
}

// NewGRPC - constructor.
func NewGRPC(downloader service.Downloader, opts ...GRPCOption) *GRPC {
	g := &GRPC{
		downloader: downloader,
	}

	for i := range opts {
		opts[i](g)
	}

	return g
}
