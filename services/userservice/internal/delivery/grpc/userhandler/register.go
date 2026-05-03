package userhandler

import (
	"context"
	"myapp/api/gen/user"
	"myapp/pkg/richerror"
	"userapp/internal/param"
)

func (h *Handler) Register(ctx context.Context,
	req *user.RegisterRequest) (*user.RegisterResponse, error) {
	const op = "userhandler.Register"

	input := param.RegisterRequest{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}

	if fieldErrors, err := h.userValidator.ValidateRegisterRequest(input); err != nil {
		return &user.RegisterResponse{}, richerror.New(op).WithMeta(map[string]interface{}{
			"validationErrors": fieldErrors,
		}).WithErr(err)
	}

	resp, err := h.userSvc.Register(ctx, input)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResponse{User: &user.UserInfo{
		ID:          uint64(resp.User.ID),
		Name:        resp.User.Name,
		PhoneNumber: resp.User.PhoneNumber,
	}}, nil

}
