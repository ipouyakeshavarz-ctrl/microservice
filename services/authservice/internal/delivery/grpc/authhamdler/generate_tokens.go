package authhandler

import (
	"context"
	"myapp/api/gen/auth"
)

func (h *Handler) GenerateTokens(ctx context.Context, req *auth.UserInfo) (*auth.LoginTokenResponse, error) {
	resp, err := h.authService.GenerateTokens(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
