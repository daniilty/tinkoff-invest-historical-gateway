package postgres

import "context"

func (d *DBImpl) GetCandlesFromTo(ctx context.Context, figi string, from uint32, to uint32) (Candles, error) {
	const q = `select id, figi, ts, hp, lp, open, close from candles
    where ts >= $1 and ts <= $2 and figi=$3 order by ts asc`

	candles := Candles{}

	rows, err := d.pool.Query(ctx, q, from, to, figi)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := Candle{}

		err = rows.Scan(&c.ID, &c.FIGI, &c.TimeStamp, &c.HighestPrice, &c.LowestPrice, &c.Open, &c.Close)
		if err != nil {
			return nil, err
		}

		candles = append(candles, c)
	}

	return candles, nil
}

func (d *DBImpl) InsertCandle(ctx context.Context, candle Candle) error {
	const q = `insert into candles(figi, ts, hp, lp, open, close)
    values($1, $2, $3, $4, $5, $6)`

	_, err := d.pool.Exec(ctx, q, candle.FIGI, candle.TimeStamp, candle.HighestPrice, candle.LowestPrice,
		candle.Open, candle.Close)

	return err
}
