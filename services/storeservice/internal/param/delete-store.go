package param

type DeleteStoreRequest struct {
	StoreID uint `json:"store_id"`
	UserID  uint `json:"user_id"`
}

type DeleteStoreResponse struct {
	Success bool `json:"success"`
}
