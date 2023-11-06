package orderbook

import "errors"

type Orderbook struct {
	BuyTree    *Limit
	SellTree   *Limit
	LowestSell *Limit
	HighestBuy *Limit
}

func NewOrderbook() *Orderbook {
	return &Orderbook{
		BuyTree:    nil,
		SellTree:   nil,
		LowestSell: nil,
		HighestBuy: nil,
	}
}

var errOrderBookNotCreated = errors.New("err: order book not created")

func (o *Orderbook) Add(order *Order) error {
	if o == nil {
		return errOrderBookNotCreated
	}

	if order.BuyOrSell {
		if o.LowestSell.LimitPrice <= order.LimitPrice {
			if err := o.Execute(order); err != nil {
				return err
			}
		} else {
			if _, err := o.BuyTree.addToBuyTree(nil, order); err != nil {
				return err
			}
		}
	} else {
		if o.HighestBuy.LimitPrice >= order.LimitPrice {
			if err := o.Execute(order); err != nil {
				return err
			}
		} else {
			if err := o.SellTree.addToSellTree(nil, order); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *Orderbook) Execute(order *Order) error {
	if o == nil {
		return errOrderBookNotCreated
	}

	if order.BuyOrSell {
		lowOrder, err := o.SellTree.removeLowest(o.LowestSell)
		if err != nil {
			return err
		}

		o.log(order, lowOrder)
	} else {
		highOrder, err := o.BuyTree.removeHighest(o.HighestBuy)
		if err != nil {
			return err
		}

		o.log(order, highOrder)
	}

	return nil
}

// TODO
func (o *Orderbook) log(order, mOrder *Order) error {
	return nil
}
