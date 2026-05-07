package param

type CheckOutRequest struct {
	UserID uint `json:"user_id"`
}

type CheckoutResponse struct {
	CheckoutID string `json:"checkout_id"`
}
