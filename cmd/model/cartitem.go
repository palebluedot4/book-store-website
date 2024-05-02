package model

type CartItem struct {
	ID     int64
	Book   *Book
	Count  int64
	Amount float64
	CartID string
}

func (c *CartItem) GetAmount() float64 {
	return c.Book.Price * float64(c.Count)
}
