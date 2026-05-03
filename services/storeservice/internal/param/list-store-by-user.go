package param

type ListStoresByUserRequest struct {
	UserID uint `json:"user_id"`
}
type ListStoresByUserResponse struct {
	Stores []StoreInfo `json:"stores"`
}
