package service

import (
	"context"
	"fmt"

	schema "github.com/daniilty/tinkoff-invest-grpc-schema"
	"github.com/daniilty/tinkoff-invest-historical-gateway/internal/data"
)

// GetCandles - get candles from db, check if requested data is missing, download it from tinkoff,
// normalize, download to db and return it to user.
func (d *DownloaderImpl) GetCandles(ctx context.Context, req CandlesRequest) (Candles, error) {
	dbCandles, err := d.db.GetCandlesFromTo(ctx, req.FIGI, req.From, req.To)
	if err != nil {
		return nil, fmt.Errorf("get db candles: %w", err)
	}

	candles := convertDBCandlesToInner(dbCandles)
	timestamps := getCandlesTimestamps(candles)

	intervals := data.GetMissingIntervals(timestamps, &data.Interval{
		Step: req.Interval,
		From: req.From,
		To:   req.To,
	})

	pbRequests := convertDataIntervalsToCandlesRequests(intervals, req.FIGI)
	pbCandles := []*schema.Candle{}

	for i := range pbRequests {
		resp, err := d.client.GetCandles(ctx, pbRequests[i])
		if err != nil {
			return nil, fmt.Errorf("get pb candles: %w", err)
		}

		pbCandles = append(pbCandles, resp.GetCandles()...)
	}

	pbCandles = normalizePBCandles(pbCandles, req.Interval)

	for i := range pbCandles {
		candles = append(candles, convertPBCandleToInner(pbCandles[i], req.FIGI))

		err = d.db.InsertCandle(ctx, convertPBCandleToDB(pbCandles[i], req.FIGI))
		if err != nil {
			return nil, fmt.Errorf("insert candle: %w", err)
		}
	}

	return candles, nil
}
