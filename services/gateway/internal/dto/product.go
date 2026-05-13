package dto

type ProductInfo struct {
	ID          uint64  `json:"id" example:"1"`
	StoreID     uint64  `json:"store_id" example:"10"`
	Name        string  `json:"name" example:"Laptop"`
	Description string  `json:"description" example:"High performance laptop"`
	Category    string  `json:"category" example:"electronics"`
	Price       float32 `json:"price" example:"2500.5"`
	Stock       int64   `json:"stock" example:"20"`
	SKU         string  `json:"sku" example:"LAPTOP-001"`
	ImageURL    string  `json:"image_url" example:"https://cdn.example.com/laptop.png"`
	IsActive    bool    `json:"is_active" example:"true"`
}

type DeleteProductRequest struct {
	StoreID   uint64 `json:"store_id"`
	UserID    uint64 `json:"user_id"`
	ProductID uint64 `json:"product_id"`
}

type DeleteProductResponse struct {
	Success bool `json:"success" example:"true"`
}

type UpdateProductRequest struct {
	ID          uint64  `json:"id"`
	StoreID     uint64  `json:"store_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	Stock       int64   `json:"stock"`
	SKU         string  `json:"sku"`
	ImageURL    string  `json:"image_url"`
	IsActive    bool    `json:"is_active"`
}

type UpdateProductResponse struct {
	Product ProductInfo `json:"product"`
}

type CreateProductRequest struct {
	StoreID     uint64  `json:"store_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	Stock       int64   `json:"stock"`
	SKU         string  `json:"sku"`
	ImageURL    string  `json:"image_url"`
	IsActive    bool    `json:"is_active"`
}

type CreateProductResponse struct {
	Product ProductInfo `json:"product"`
}

type GetProductResponse struct {
	Product ProductInfo `json:"product"`
}

type GetProductByIDRequest struct {
	ProductID uint64 `json:"product_id" example:"1"`
	StoreID   uint64 `json:"store_id" example:"10"`
}
