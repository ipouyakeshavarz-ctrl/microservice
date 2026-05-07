package authhandler

import (
	"context"
	"myapp/api/gen/auth"
)

func (h *Handler) VerifyToken(ctx context.Context, req *auth.VerifyTokenRequest) (*auth.VerifyTokenResponse, error) {

	resp, err := h.authService.VerifyToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}
