package orderbook

// orders are represented as linked-lists in a tree of Limit objects
type Order struct {
	ID          int
	BuyOrSell   bool
	Quantity    int
	LimitPrice  float32
	EntryTime   int
	EventTime   int
	next        *Order
	prev        *Order
	parentLimit *Limit
}

func NewOrder(id, quantity, entryTime, eventTIme int, buyOrSell bool, limit float32) *Order {
	return &Order{
		ID:          id,
		Quantity:    quantity,
		EntryTime:   entryTime,
		EventTime:   eventTIme,
		BuyOrSell:   buyOrSell,
		LimitPrice:  limit,
		next:        nil,
		prev:        nil,
		parentLimit: nil,
	}
}

func (o *Order) Next() *Order {
	return o.next
}

func (o *Order) Prev() *Order {
	return o.prev
}

func (o *Order) GetParentLimit() *Limit {
	return o.parentLimit
}
