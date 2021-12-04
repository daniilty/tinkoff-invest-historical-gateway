package main

import (
	"fmt"
	"os"
)

type envConfig struct {
	grpcAddr               string
	tinkoffGatewayGrpcAddr string
	pgConnString           string
}

func loadEnvConfig() (*envConfig, error) {
	const (
		provideEnvErrorMsg = `please provide "%s" environment variable`

		grpcAddrEnv               = "GRPC_SERVER_ADDR"
		tinkoffGatewayGrpcAddrEnv = "TINKOFF_GATEWAY_GRPC_ADDR"
		pgConnStringEnv           = "PG_CONNSTRING"
	)

	var ok bool

	cfg := &envConfig{}

	cfg.grpcAddr, ok = os.LookupEnv(grpcAddrEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, grpcAddrEnv)
	}

	cfg.tinkoffGatewayGrpcAddr, ok = os.LookupEnv(tinkoffGatewayGrpcAddrEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, tinkoffGatewayGrpcAddrEnv)
	}

	cfg.pgConnString, ok = os.LookupEnv(pgConnStringEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, pgConnStringEnv)
	}

	return cfg, nil
}
