package dto

type CartItem struct {
	ProductID uint64 `json:"product_id" example:"10"`
	Quantity  int32  `json:"quantity" example:"2"`
}

type Cart struct {
	UserID uint64     `json:"user_id" example:"5"`
	Items  []CartItem `json:"items"`
}

type AddItemRequest struct {
	UserID    uint64 `json:"user_id" example:"5"`
	ProductID uint64 `json:"product_id" example:"10"`
	Quantity  int32  `json:"quantity" example:"2"`
}

type AddItemResponse struct {
	Success bool `json:"success" example:"true"`
}

type GetCartRequest struct {
	UserID uint64 `json:"user_id" example:"5"`
}

type GetCartResponse struct {
	Cart Cart `json:"cart"`
}

type RemoveItemRequest struct {
	UserID    uint64 `json:"user_id" example:"5"`
	ProductID int64  `json:"product_id" example:"10"`
}

type RemoveItemResponse struct {
	Success bool `json:"success" example:"true"`
}

type CheckoutRequest struct {
	UserID uint64 `json:"user_id" example:"5"`
}

type CheckoutResponse struct {
	CheckoutID string `json:"checkout_id" example:"checkout_123"`
}
