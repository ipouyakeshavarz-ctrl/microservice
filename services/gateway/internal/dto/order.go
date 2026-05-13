package dto

type OrderItem struct {
	ProductID uint64 `json:"product_id" example:"10"`
	Quantity  int32  `json:"quantity" example:"2"`
	Price     uint64 `json:"price" example:"50000"`
}

type Order struct {
	ID         string      `json:"id" example:"order_123"`
	CheckoutID string      `json:"checkout_id" example:"checkout_123"`
	UserID     uint64      `json:"user_id" example:"5"`
	Items      []OrderItem `json:"items"`
	TotalPrice uint64      `json:"total_price" example:"100000"`
	Status     string      `json:"status" example:"paid"`
	CreatedAt  int64       `json:"created_at" example:"1710000000"`
}

type GetOrderRequest struct {
	UserID  uint64 `json:"user_id" example:"5"`
	OrderID string `json:"order_id" example:"order_123"`
}

type GetOrderResponse struct {
	Order Order `json:"order"`
}

type ListUserOrdersRequest struct {
	UserID uint64 `json:"user_id" example:"5"`
}

type ListUserOrdersResponse struct {
	Orders []Order `json:"orders"`
}
