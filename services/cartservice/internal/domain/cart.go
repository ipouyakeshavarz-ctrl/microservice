package domain

import "time"

type Cart struct {
	UserID    uint
	Items     []CartItem
	UpdatedAt time.Time
}

type CartItem struct {
	ProductID uint
	Quantity  uint
}

func NewCart(userID uint, now time.Time) *Cart {
	return &Cart{
		UserID:    userID,
		Items:     []CartItem{},
		UpdatedAt: now,
	}
}

func (c *Cart) IsEmpty() bool {
	return len(c.Items) == 0
}

func (c *Cart) AddItem(productID uint, qty uint, now time.Time) {
	for i := range c.Items {
		if c.Items[i].ProductID == productID {
			c.Items[i].Quantity += qty
			c.UpdatedAt = now
			return
		}
	}

	c.Items = append(c.Items, CartItem{
		ProductID: productID,
		Quantity:  qty,
	})
	c.UpdatedAt = now
}

func (c *Cart) RemoveItem(productID uint, now time.Time) bool {
	for i := range c.Items {
		if c.Items[i].ProductID == productID {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			c.UpdatedAt = now
			return true
		}
	}
	return false
}

func (c *Cart) UpdateQuantity(productID uint, qty uint, now time.Time) bool {
	for i := range c.Items {
		if c.Items[i].ProductID == productID {
			c.Items[i].Quantity = qty
			c.UpdatedAt = now
			return true
		}
	}
	return false
}
