package orderbook

type Limit struct {
	LimitPrice  float32
	Size        int
	TotalVolume int
	Parent      *Limit
	LeftChild   *Limit
	RightChild  *Limit
	HeadOrder   *Order
	TailOrder   *Order
}
