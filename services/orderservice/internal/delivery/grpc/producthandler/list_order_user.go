package orderhandler

import (
	"context"
	"strconv"

	"myapp/api/gen/order"
)

func (h Handler) ListUserOrders(
	ctx context.Context,
	req *order.ListUserOrdersRequest,
) (*order.ListUserOrdersResponse, error) {

	orders, err := h.orderSvc.ListUserOrders(ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}

	resp := &order.ListUserOrdersResponse{
		Orders: make([]*order.Order, 0, len(orders)),
	}

	for _, o := range orders {
		items := make([]*order.OrderItem, 0, len(o.Items))

		for _, item := range o.Items {
			items = append(items, &order.OrderItem{
				ProductId: uint64(item.ProductId),
				Quantity:  int32(item.Quantity),
			})
		}

		resp.Orders = append(resp.Orders, &order.Order{
			Id:         strconv.Itoa(int(o.Id)),
			CheckoutId: o.CheckoutId,
			UserId:     uint64(o.UserId),
			Status:     o.Status,
			Items:      items,
			CreatedAt:  o.CreatedAt,
		})
	}

	return resp, nil
}
