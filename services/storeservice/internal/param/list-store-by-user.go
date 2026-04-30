package param

import "storeapp/internal/entity"

type ListStoresByUserRequest struct {
	UserID uint `json:"user_id"`
}
type ListStoresByUserResponse struct {
	Stores []entity.Store `json:"stores"`
}
