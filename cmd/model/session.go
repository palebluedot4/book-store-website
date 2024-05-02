package model

type Session struct {
	SessionID string
	Username  string
	UserID    int64
	Cart      *Cart
	Order     *Order
	Orders    []*Order
}
