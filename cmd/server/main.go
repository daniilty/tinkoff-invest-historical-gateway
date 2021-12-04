package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	schema "github.com/daniilty/tinkoff-invest-grpc-schema"
	"github.com/daniilty/tinkoff-invest-historical-gateway/internal/postgres"
	"github.com/daniilty/tinkoff-invest-historical-gateway/internal/server"
	"github.com/daniilty/tinkoff-invest-historical-gateway/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func run() error {
	cfg, err := loadEnvConfig()
	if err != nil {
		return err
	}

	pool, err := pgxpool.Connect(context.Background(), cfg.pgConnString)
	if err != nil {
		return err
	}

	conn, err := grpc.DialContext(context.Background(), cfg.tinkoffGatewayGrpcAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	investClient := schema.NewTinkoffInvestAPIGatewayClient(conn)

	db := postgres.NewDBImpl(pool)
	downloader := service.NewDownloaderImpl(db, investClient)

	loggerCfg := zap.NewProductionConfig()

	logger, err := loggerCfg.Build()
	if err != nil {
		return err
	}

	sugaredLogger := logger.Sugar()

	historicalGRPCServer := server.NewGRPC(downloader, server.WithZapLogger(sugaredLogger))

	grpcServer := grpc.NewServer()
	schema.RegisterTinkoffInvestHistoricalGatewayServer(grpcServer, historicalGRPCServer)

	listener, err := net.Listen("tcp", cfg.grpcAddr)
	if err != nil {
		return err
	}

	sugaredLogger.Infow("GRPC server is starting.", "addr", listener.Addr())
	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			sugaredLogger.Errorw("Server failed to start.", "err", err)
		}
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-termChan

	sugaredLogger.Info("Gracefully stopping GRPC server.")
	grpcServer.GracefulStop()

	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
