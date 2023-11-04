package orderbook

type Orderbook struct {
	BuyTree    *Limit
	SellTree   *Limit
	LowestSell *Limit
	HighestBuy *Limit
}
