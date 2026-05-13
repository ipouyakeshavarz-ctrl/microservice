package mapper

import (
	"gatewayapp/internal/dto"
	"myapp/api/gen/cart"
)

func ToCartItem(item *cart.CartItem) dto.CartItem {
	if item == nil {
		return dto.CartItem{}
	}

	return dto.CartItem{
		ProductID: item.ProductId,
		Quantity:  item.Quantity,
	}
}

func ToCart(cart *cart.Cart) dto.Cart {
	if cart == nil {
		return dto.Cart{}
	}

	items := make([]dto.CartItem, 0, len(cart.Items))

	for _, i := range cart.Items {
		items = append(items, ToCartItem(i))
	}

	return dto.Cart{
		UserID: cart.UserId,
		Items:  items,
	}
}

func ToAddItemResponse(_ *cart.AddItemResponse) *dto.AddItemResponse {
	return &dto.AddItemResponse{
		Success: true,
	}
}

func ToGetCartResponse(res *cart.GetCartResponse) *dto.GetCartResponse {
	if res == nil {
		return &dto.GetCartResponse{}
	}

	return &dto.GetCartResponse{
		Cart: ToCart(res.Cart),
	}
}

func ToCheckoutResponse(res *cart.CheckoutResponse) *dto.CheckoutResponse {
	if res == nil {
		return &dto.CheckoutResponse{}
	}

	return &dto.CheckoutResponse{
		CheckoutID: res.CheckoutId,
	}
}

func ToAddItemRequest(req *dto.AddItemRequest) *cart.AddItemRequest {
	return &cart.AddItemRequest{
		UserId:    req.UserID,
		ProductId: req.ProductID,
		Quantity:  req.Quantity,
	}
}

func ToGetCartRequest(req *dto.GetCartRequest) *cart.GetCartRequest {
	return &cart.GetCartRequest{
		UserId: req.UserID,
	}
}

func ToCheckoutRequest(req *dto.CheckoutRequest) *cart.CheckoutRequest {
	return &cart.CheckoutRequest{
		UserId: req.UserID,
	}
}
