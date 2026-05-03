package domain

import "time"

type Product struct {
	ID          uint
	StoreID     uint
	Name        string
	Description string
	Category    Category
	Price       float64
	Stock       int
	SKU         string
	ImageURL    string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
