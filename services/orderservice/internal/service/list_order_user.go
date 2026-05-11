package orderservice

import (
	"context"
	"orderapp/internal/param"
)

func (s *Service) ListUserOrders(
	ctx context.Context,
	userID uint,
) ([]param.OrderInfo, error) {

	orders, err := s.repo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	resp := make([]param.OrderInfo, 0, len(orders))

	for _, o := range orders {
		items := make([]param.OrderItem, 0, len(o.Items))

		for _, item := range o.Items {
			items = append(items, param.OrderItem{
				ProductId: item.ProductID,
				Quantity:  item.Quantity,
			})
		}

		resp = append(resp, param.OrderInfo{
			Id:         o.ID,
			CheckoutId: o.CheckoutID,
			UserId:     o.UserID,
			Status:     string(o.Status),
			Items:      items,
			CreatedAt:  o.CreatedAt.Unix(),
		})
	}

	return resp, nil
}
