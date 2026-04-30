package param

type GetProductByIDRequest struct {
	ID      uint `json:"id"`
	StoreID uint `json:"store_id"`
}

type GetProductResponse struct {
	Product ProductInfo `json:"product"`
}
