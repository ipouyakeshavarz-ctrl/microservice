package userservice

import (
	"context"
	"myapp/pkg/richerror"
	"userapp/internal/param"
)

func (s Service) Profile(ctx context.Context, req param.ProfileRequest) (param.ProfileResponse, error) {
	const op = "userservice.Profile"

	user, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return param.ProfileResponse{}, richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"req": req})
	}

	return param.ProfileResponse{ID: user.ID,
		Name: user.Name,
	}, nil
}
