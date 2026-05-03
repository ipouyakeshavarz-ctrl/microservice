package userhandler

import (
	"context"
	"myapp/api/gen/user"
	"myapp/pkg/richerror"

	"userapp/internal/param"
)

func (h *Handler) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	const op = "userHandlergrpc.Login"
	input := param.LoginRequest{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}

	if fieldErrors, err := h.userValidator.ValidateLoginRequest(input); err != nil {
		return &user.LoginResponse{}, richerror.New(op).WithMeta(map[string]interface{}{
			"validationErrors": fieldErrors,
		}).WithErr(err)
	}

	resp, err := h.userSvc.Login(ctx, input)
	if err != nil {
		return nil, err
	}

	return &user.LoginResponse{
		User: &user.UserInfo{
			ID:          uint64(resp.User.ID),
			PhoneNumber: resp.User.PhoneNumber,
			Name:        resp.User.Name,
		},
		Tokens: &user.Tokens{
			AccessToken:  resp.Tokens.AccessToken,
			RefreshToken: resp.Tokens.RefreshToken,
		}}, nil
}
