package model

type Order struct {
	OrderID     string
	CreateTime  string
	TotalCount  int64
	TotalAmount float64
	Status      OrderStatus
	UserID      int64
}

type OrderStatus int64

const (
	Pending OrderStatus = iota
	Shipped
	Received
)

func (order *Order) PendingShipping() bool {
	return order.Status == 0
}
func (order *Order) OrderShipped() bool {
	return order.Status == 1
}
func (order *Order) OrderReceived() bool {
	return order.Status == 2
}
