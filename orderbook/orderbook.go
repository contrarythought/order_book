package orderbook

import "time"

type Order struct {
	Price     float32
	Quantity  uint
	TimeStamp time.Time
	Action    byte
}

type Orderbook struct {
}
