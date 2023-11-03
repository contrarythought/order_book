package orderbook

import (
	"errors"
	"time"
)

type Order struct {
	Price      float32
	Quantity   uint
	TimeStamp  time.Time
	Action     byte
	samePrices []*Order
	right      *Order
	left       *Order
	parent     *Order
}

func NewOrder(price float32, quantity uint, timeStamp time.Time, action byte) *Order {
	return &Order{
		Price:      price,
		Quantity:   quantity,
		TimeStamp:  timeStamp,
		Action:     action,
		samePrices: nil,
		right:      nil,
		left:       nil,
		parent:     nil,
	}
}

type Orderbook struct {
	sellRoot *Order // root of heap that stores all of the sell orders
	buyRoot  *Order // root of heap that stores all of the buy orders
}

var errOrderbookNotCreated = errors.New("err: order book not created")

func NewOrderBook() *Orderbook {
	return &Orderbook{
		sellRoot: nil,
		buyRoot:  nil,
	}
}

func (o *Orderbook) MatchBuyOrderWithSell(buyOrder *Order) (*Order, error) {
	if o != nil {
		if o.sellRoot.Price > buyOrder.Price {
			if err := o.InsertOrder(buyOrder); err != nil {
				return nil, err
			}
			return nil, nil
		}
		return o.sellRoot, nil
	}

	return nil, errOrderbookNotCreated
}

func (o *Orderbook) MatchSellOrderWithBuy(sellOrder *Order) (*Order, error) {
	if o != nil {
		if o.buyRoot.Price < sellOrder.Price {
			if err := o.InsertOrder(sellOrder); err != nil {
				return nil, err
			}
			return nil, nil
		}
		return o.buyRoot, nil
	}

	return nil, errOrderbookNotCreated
}

const (
	BUY  = 1
	SELL = 0
)

func (o *Orderbook) InsertOrder(order *Order) error {
	if o == nil {
		return errOrderbookNotCreated
	}

	switch order.Action {
	case BUY:
		if err := o.insertIntoBuyHeap(order); err != nil {
			return err
		}
	case SELL:
		if err := o.insertIntoSellHeap(order); err != nil {
			return err
		}
	default:
		return errors.New("err: unrecognized action")
	}

	return nil
}

// max heap
func (o *Orderbook) insertIntoBuyHeap(order *Order) error {
	if o == nil {
		return errOrderbookNotCreated
	}

	// if the order book has another order with the same price as the incoming order
	// add the incoming order into the queue of that price level
	sp, err := o.searchBuyHeap(order.Price)
	if err != nil {
		return err
	}
	if sp != nil {
		sp.samePrices = append(sp.samePrices, order)
		return nil
	}

	if o.buyRoot == nil {
		o.buyRoot = NewOrder(order.Price, order.Quantity, order.TimeStamp, order.Action)
	} else {

	}

	return nil
}

func (o *Orderbook) searchBuyHeap(price float32) (*Order, error) {
	return nil, nil
}

func (o *Orderbook) insertIntoSellHeap(order *Order) error {
	if o == nil {
		return errOrderbookNotCreated
	}

	return nil
}

func (o *Orderbook) searchSellHeap(price float32) error {
	return nil
}
