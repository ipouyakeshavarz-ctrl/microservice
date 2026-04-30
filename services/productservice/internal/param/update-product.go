package param

import "productapp/internal/entity"

type UpdateProductRequest struct {
	ID          uint            `json:"id"`
	StoreID     uint            `json:"store_id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Category    entity.Category `json:"category"`
	Price       float64         `json:"price"`
	Stock       int             `json:"stock"`
	SKU         string          `json:"sku"`
	ImageURL    string          `json:"image_url"`
	IsActive    bool            `json:"is_active"`
}

type UpdateProductResponse struct {
	Product ProductInfo `json:"product"`
}
