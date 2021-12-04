package server

import "go.uber.org/zap"

// GRPCOption - configurable gateway options.
type GRPCOption func(*GRPC)

// WithToken - set here tinkoff api token.
func WithZapLogger(logger *zap.SugaredLogger) GRPCOption {
	return func(g *GRPC) {
		g.logger = logger
	}
}
