package mapper

import (
	"gatewayapp/internal/dto"
	"myapp/api/gen/order"
)

func ToOrderItem(item *order.OrderItem) dto.OrderItem {
	if item == nil {
		return dto.OrderItem{}
	}

	return dto.OrderItem{
		ProductID: item.ProductId,
		Quantity:  item.Quantity,
		Price:     item.Price,
	}
}

func ToOrder(order *order.Order) dto.Order {
	if order == nil {
		return dto.Order{}
	}

	items := make([]dto.OrderItem, 0, len(order.Items))
	for _, i := range order.Items {
		items = append(items, ToOrderItem(i))
	}

	return dto.Order{
		ID:         order.Id,
		CheckoutID: order.CheckoutId,
		UserID:     order.UserId,
		Items:      items,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
	}
}

func ToGetOrderRequest(req *dto.GetOrderRequest) *order.GetOrderRequest {
	return &order.GetOrderRequest{
		UserId:  req.UserID,
		OrderId: req.OrderID,
	}
}

func ToGetOrderResponse(res *order.GetOrderResponse) *dto.GetOrderResponse {
	if res == nil {
		return &dto.GetOrderResponse{}
	}

	return &dto.GetOrderResponse{
		Order: ToOrder(res.Order),
	}
}

func ToListUserOrdersRequest(req *dto.ListUserOrdersRequest) *order.ListUserOrdersRequest {
	return &order.ListUserOrdersRequest{
		UserId: req.UserID,
	}
}

func ToListUserOrdersResponse(res *order.ListUserOrdersResponse) *dto.ListUserOrdersResponse {
	if res == nil {
		return &dto.ListUserOrdersResponse{}
	}

	orders := make([]dto.Order, 0, len(res.Orders))

	for _, o := range res.Orders {
		orders = append(orders, ToOrder(o))
	}

	return &dto.ListUserOrdersResponse{
		Orders: orders,
	}
}
