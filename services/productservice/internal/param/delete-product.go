package param

type DeleteProductRequest struct {
	ID      uint `json:"id"`
	StoreID uint `json:"store_id"`
	UserID  uint `json:"user_id"`
}

type DeleteProductResponse struct {
	Success bool `json:"success"`
}
