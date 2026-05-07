package param

type GetCartRequest struct {
	UserID uint `json:"user_id"`
}
type GetCartResponse struct {
	UserID uint
	Items  []CartItemView
}

type CartItemView struct {
	ProductID uint
	Quantity  uint
}
