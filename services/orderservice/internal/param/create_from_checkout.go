package param

import "orderapp/internal/domain"

type CreateFromCheckoutRequest struct {
	Event domain.CartCheckedOutEvent `json:"event"`
}
