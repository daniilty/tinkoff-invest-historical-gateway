package service

type CandlesRequest struct {
	Interval uint32
	From     uint32
	To       uint32
	FIGI     string
}
