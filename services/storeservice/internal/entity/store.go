package entity

type Store struct {
	ID          uint
	UserID      uint
	Name        string
	Description string
	LogoURL     string
	Address     Address
	PhoneNumber string
	IsActive    bool
}

type Address struct {
	Street      string
	City        string
	Province    string
	PostalCode  string
	Description string
}
