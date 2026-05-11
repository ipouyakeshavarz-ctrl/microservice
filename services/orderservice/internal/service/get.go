package orderservice

import (
	"context"
	"orderapp/internal/param"
)

func (s *Service) GetOrder(
	ctx context.Context,
	req param.GetOrderRequest,
) (param.GetOrderResponse, error) {

	o, err := s.repo.GetByID(ctx, uint(req.OrderId))
	if err != nil {
		return param.GetOrderResponse{}, err
	}

	items := make([]param.OrderItem, 0, len(o.Items))

	for _, item := range o.Items {
		items = append(items, param.OrderItem{
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return param.GetOrderResponse{
		Order: param.OrderInfo{
			Id:         o.ID,
			CheckoutId: o.CheckoutID,
			UserId:     o.UserID,
			Status:     string(o.Status),
			Items:      items,
			CreatedAt:  o.CreatedAt.Unix(),
		},
	}, nil
}
