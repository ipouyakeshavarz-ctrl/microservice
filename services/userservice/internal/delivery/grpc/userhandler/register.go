package userhandler

import (
	"context"
	"myapp/api/gen/user"
	"myapp/pkg/errmsg"
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
		return nil, richerror.New(op).
			WithKind(richerror.KindInvalid).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithFields(fieldErrors)
	}

	resp, err := h.userSvc.Register(ctx, input)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithFields(map[string]string{
			"massage": err.Error(),
		})
	}

	return &user.RegisterResponse{User: &user.UserInfo{
		ID:          uint64(resp.User.ID),
		Name:        resp.User.Name,
		PhoneNumber: resp.User.PhoneNumber,
	}}, nil

}
