package mapper

import (
	"gatewayapp/internal/dto"
	"myapp/api/gen/user"
)

func ToLoginRequest(req *dto.LoginRequest) *user.LoginRequest {
	return &user.LoginRequest{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}
}

func ToLoginResponse(resp *user.LoginResponse) *dto.LoginResponse {
	return &dto.LoginResponse{
		User:   ToUserInfo(resp.User),
		Tokens: ToTokens(resp.Tokens),
	}

}

func ToUserInfo(req *user.UserInfo) dto.UserInfo {
	return dto.UserInfo{
		ID:          req.ID,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}
}

func ToTokens(req *user.Tokens) dto.Tokens {
	return dto.Tokens{
		AccessToken:  req.AccessToken,
		RefreshToken: req.RefreshToken,
	}

}

func ToProfileRequest(req *dto.ProfileRequest) *user.ProfileRequest {
	return &user.ProfileRequest{
		UserId: req.UserId,
	}
}

func ToProfileResponse(resp *user.ProfileResponse) *dto.ProfileResponse {
	return &dto.ProfileResponse{
		UserID: resp.UserId,
		Name:   resp.Name,
	}
}

func ToRegisterRequest(req *dto.RegisterRequest) *user.RegisterRequest {
	return &user.RegisterRequest{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
		Name:        req.Name,
	}
}

func ToRegisterResponse(resp *user.RegisterResponse) *dto.RegisterResponse {
	return &dto.RegisterResponse{
		User: ToUserInfo(resp.User),
	}
}
