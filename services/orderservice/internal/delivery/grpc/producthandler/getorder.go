package orderhandler

import (
	"context"
	"myapp/api/gen/order"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"orderapp/internal/param"
	"strconv"
)

func (h Handler) GetOrder(
	ctx context.Context,
	req *order.GetOrderRequest,
) (*order.GetOrderResponse, error) {
	const op = "orderHandler.GetOrder"

	orderID, err := strconv.Atoi(req.OrderId)
	if err != nil {
		return nil, richerror.New(op).
			WithKind(richerror.KindInvalid).
			WithMessage(errmsg.ErrorMsgInvalidInput)
	}

	resp, gErr := h.orderSvc.GetOrder(ctx, param.GetOrderRequest{
		OrderId: orderID,
	})

	if &resp != nil {
		if uint(req.UserId) != resp.Order.UserId {
			return nil, richerror.New(op).
				WithKind(richerror.KindForbidden).
				WithMessage(errmsg.ErrorMsgUserNotAllowed)
		}
	}

	if gErr != nil {
		return nil, richerror.New(op).WithErr(gErr)
	}

	items := make([]*order.OrderItem, 0, len(resp.Order.Items))

	for _, item := range resp.Order.Items {
		items = append(items, &order.OrderItem{
			ProductId: uint64(item.ProductId),
			Quantity:  int32(item.Quantity),
		})
	}

	return &order.GetOrderResponse{
		Order: &order.Order{
			Id:         strconv.Itoa(int(resp.Order.Id)),
			CheckoutId: resp.Order.CheckoutId,
			UserId:     uint64(resp.Order.UserId),
			Status:     resp.Order.Status,
			Items:      items,
			CreatedAt:  resp.Order.CreatedAt,
		},
	}, nil
}
