package param

import (
	"storeapp/internal/domain"
)

type StoreInfo struct {
	ID          uint           `json:"id"`
	UserID      uint           `json:"user_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	LogoURL     string         `json:"logo_url"`
	Address     domain.Address `json:"address"`
	PhoneNumber string         `json:"phone_number"`
	IsActive    bool           `json:"is_active"`
}
