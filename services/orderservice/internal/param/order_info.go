package param

type OrderItem struct {
	ProductId uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type OrderInfo struct {
	Id         uint        `json:"id"`
	CheckoutId string      `json:"checkout_id"`
	UserId     uint        `json:"user_id"`
	Status     string      `json:"status"`
	Items      []OrderItem `json:"items"`
	CreatedAt  int64       `json:"created_at"`
}
