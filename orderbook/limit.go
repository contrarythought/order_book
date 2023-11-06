package orderbook

import "errors"

type Limit struct {
	LimitPrice  float32
	Size        int
	TotalVolume int
	Parent      *Limit
	LeftChild   *Limit
	RightChild  *Limit
	Orders      []*Order
}

var errLimitNotCreated = errors.New("err: limit node not created")

func NewLimit(parent *Limit, order *Order) *Limit {
	return &Limit{
		LimitPrice:  order.LimitPrice,
		Size:        order.Quantity, // total outstanding live order quantities
		TotalVolume: order.Quantity, // total quantity of current and past orderse
		Parent:      parent,
		Orders:      []*Order{order},
		LeftChild:   nil,
		RightChild:  nil,
	}
}

// TODO: rethink this
func (l *Limit) addToBuyTree(root *Limit, order *Order) (*Limit, error) {
	if l == nil {
		l = NewLimit(root, order)
		return l, nil
	}

	if order.LimitPrice < l.LimitPrice {
		l.addToBuyTree(l.LeftChild, order)
	} else if order.LimitPrice > l.LimitPrice {
		l.addToBuyTree(l.RightChild, order)
	} else {
		l.Orders = append(l.Orders, order)
		l.Size += order.Quantity
		l.TotalVolume += order.Quantity
	}

	return l, nil
}

func (l *Limit) addToSellTree(root *Limit, order *Order) error {
	return nil
}

func (l *Limit) removeLowest(lowestSell *Limit) (*Order, error) {
	if l == nil {
		return nil, errLimitNotCreated
	}

	var o *Order
	if len(l.Orders) == 0 {
		lowestSell.Parent.LeftChild = nil
	} else {
		o = l.Orders[0]
		l.Orders = l.Orders[1:]
	}
	return o, nil
}

func (l *Limit) removeHighest(highestBuy *Limit) (*Order, error) {
	if l == nil {
		return nil, errLimitNotCreated
	}

	var o *Order
	if len(l.Orders) == 0 {
		highestBuy.Parent.RightChild = nil
	} else {
		o = highestBuy.Orders[0]
		highestBuy.Orders = highestBuy.Orders[1:]
	}

	return o, nil
}
