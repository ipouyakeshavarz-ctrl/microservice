package userservice

import (
	"context"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"userapp/internal/param"
)

func (s Service) Login(ctx context.Context, req param.LoginRequest) (param.LoginResponse, error) {
	const op = "userservice.Login"

	// TODO - it would be better to param two separate method for existence check and getUserByPhoneNumber
	user, err := s.repo.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		return param.LoginResponse{}, richerror.New(op).WithErr(err).
			WithMeta(map[string]interface{}{"phone_number": req.PhoneNumber})
	}

	if user.Password != getMD5Hash(req.Password) {
		return param.LoginResponse{}, richerror.New(op).
			WithKind(richerror.KindInvalid).WithMessage(errmsg.ErrorMsgUserNameOrPasswordNotCorrect)
	}

	accessToken, refreshToken, gErr := s.authClient.GenerateTokens(ctx, user)

	if gErr != nil {
		return param.LoginResponse{}, richerror.New(op).WithErr(gErr)
	}

	return param.LoginResponse{
		User: param.UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
		},
		Tokens: param.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken},
	}, nil
}
