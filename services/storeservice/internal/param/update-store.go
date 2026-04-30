package param

import (
	"storeapp/internal/entity"
)

type UpdateStoreRequest struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	LogoURL     string         `json:"logo_url"`
	Address     entity.Address `json:"address"`
	PhoneNumber string         `json:"phone_number"`
	IsActive    bool           `json:"is_active"`
}
type UpdateStoreResponse struct {
	Store StoreInfo `json:"updates-store"`
}
