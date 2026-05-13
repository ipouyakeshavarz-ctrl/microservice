package dto

type Address struct {
	Street      string `json:"street" example:"Valiasr"`
	City        string `json:"city" example:"Tehran"`
	Province    string `json:"province" example:"Tehran"`
	PostalCode  string `json:"postal_code" example:"1234567890"`
	Description string `json:"description" example:"near metro"`
}

type StoreInfo struct {
	ID          uint64  `json:"id" example:"1"`
	UserID      uint64  `json:"user_id" example:"10"`
	Name        string  `json:"name" example:"My Store"`
	Description string  `json:"description" example:"Best grocery store"`
	LogoURL     string  `json:"logo_url" example:"https://cdn.example.com/logo.png"`
	Address     Address `json:"address"`
	PhoneNumber string  `json:"phone_number" example:"09121234567"`
	IsActive    bool    `json:"is_active" example:"true"`
}

type CreateStoreRequest struct {
	UserId      uint64  `json:"user_id"`
	Name        string  `json:"name" example:"My Store"`
	Description string  `json:"description" example:"Best grocery store"`
	LogoURL     string  `json:"logo_url" example:"https://cdn.example.com/logo.png"`
	Address     Address `json:"address"`
	PhoneNumber string  `json:"phone_number" example:"09121234567"`
}

type CreateStoreResponse struct {
	Store StoreInfo `json:"store"`
}

type DeleteStoreRequest struct {
	StoreID uint64 `json:"store_id" example:"1"`
	UserID  uint64 `json:"user_id" example:"10"`
}

type DeleteStoreResponse struct {
	Success bool `json:"success" example:"true"`
}

type ListStoresByUserRequest struct {
	UserID uint64 `json:"user_id" example:"10"`
}

type ListStoresByUserResponse struct {
	Stores []StoreInfo `json:"stores"`
}

type UpdateStoreRequest struct {
	StoreID     uint64  `json:"store_id" example:"1"`
	UserID      uint64  `json:"user_id" example:"10"`
	Name        string  `json:"name" example:"Updated Store"`
	Description string  `json:"description" example:"Updated description"`
	LogoURL     string  `json:"logo_url" example:"https://cdn.example.com/logo.png"`
	Address     Address `json:"address"`
	PhoneNumber string  `json:"phone_number" example:"09121234567"`
	IsActive    bool    `json:"is_active" example:"true"`
}

type UpdateStoreResponse struct {
	Store StoreInfo `json:"store"`
}
