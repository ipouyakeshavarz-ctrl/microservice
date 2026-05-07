package domain

import "time"

type Store struct {
	ID          uint
	UserID      uint
	Name        string
	Description string
	LogoURL     string
	Address     Address
	PhoneNumber string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type Address struct {
	Street      string
	City        string
	Province    string
	PostalCode  string
	Description string
}
