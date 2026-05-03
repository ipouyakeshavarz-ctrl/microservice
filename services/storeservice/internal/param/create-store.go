package param

import "storeapp/internal/domain"

type CreateStoreRequest struct {
	UserID      uint           `json:"user_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	LogoURL     string         `json:"logo_url"`
	Address     domain.Address `json:"address"`
	PhoneNumber string         `json:"phone_number"`
}

type CreateStoreResponse struct {
	Store StoreInfo `json:"store"`
}
