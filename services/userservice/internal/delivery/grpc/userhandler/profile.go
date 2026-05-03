package userhandler

import (
	"context"
	"myapp/api/gen/user"
	"strconv"

	"userapp/internal/param"
)

func (h *Handler) GetProfile(ctx context.Context, req *user.ProfileRequest) (*user.ProfileResponse, error) {
	userid, err := strconv.Atoi(req.UserId)
	if err != nil {
		return nil, err
	}
	resp, err := h.userSvc.Profile(ctx,
		param.ProfileRequest{
			UserID: uint(userid),
		})
	if err != nil {
		return nil, err
	}

	return &user.ProfileResponse{
		UserId: strconv.Itoa(int(resp.ID)),
		Name:   resp.Name,
	}, nil
}
