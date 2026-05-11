package domain

import "time"

type OrderStatus string

const (
	OrderStatusCreated OrderStatus = "ORDER_STATUS_CREATED"
)

type Order struct {
	ID         uint
	CheckoutID string
	UserID     uint
	Status     OrderStatus
	Items      []OrderItem
	CreatedAt  time.Time
}

type OrderItem struct {
	ProductID uint
	Quantity  uint
}
