package model

type Cart struct {
	CartID      string
	CartItems   []*CartItem
	TotalCount  int64
	TotalAmount float64
	UserID      int64
	Username    string
}

func (c *Cart) GetTotalCount() int64 {
	var totalCount int64

	for _, v := range c.CartItems {
		totalCount += v.Count
	}
	return totalCount
}

func (c *Cart) GetTotalAmount() float64 {
	var totalAmount float64

	for _, v := range c.CartItems {
		totalAmount += v.GetAmount()
	}
	return totalAmount
}
