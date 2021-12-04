CREATE TABLE IF NOT EXISTS candles (
  id serial PRIMARY KEY,
  figi varchar(12) NOT NULL,
  ts bigint UNIQUE NOT NULL,
  hp double precision NOT NULL,
  lp double precision NOT NULL,
  open double precision NOT NULL,
  close double precision NOT NULL
);

