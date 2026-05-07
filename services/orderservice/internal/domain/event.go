package domain

import "time"

type CartCheckedOutEvent struct {
	CheckoutID string           `json:"checkout_id"`
	UserID     uint             `json:"user_id"`
	Items      []CheckedOutItem `json:"items"`
	OccurredAt time.Time        `json:"occurred_at"`
}

type CheckedOutItem struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
