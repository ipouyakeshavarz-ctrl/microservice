package userservice

import (
	"context"
	"myapp/pkg/richerror"
	"userapp/internal/domain"
	"userapp/internal/param"
)

func (s Service) Register(ctx context.Context, req param.RegisterRequest) (param.RegisterResponse, error) {
	const op = "userservice.Register"
	// TODO - we should verify phone number by verification code
	// TODO - replace md5 with bcrypt
	user := domain.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    getMD5Hash(req.Password),
		Role:        domain.UserRole,
	}

	createdUser, err := s.repo.Register(ctx, user)
	if err != nil {
		return param.RegisterResponse{}, richerror.New(op).WithErr(err)
	}

	return param.RegisterResponse{User: param.UserInfo{
		ID:          createdUser.ID,
		PhoneNumber: createdUser.Name,
		Name:        createdUser.PhoneNumber,
	}}, nil
}
