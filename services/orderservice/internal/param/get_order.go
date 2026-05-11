package param

type GetOrderRequest struct {
	OrderId int `json:"order_id"`
}

type GetOrderResponse struct {
	Order OrderInfo `json:"order"`
}
