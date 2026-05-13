package dto

type UserInfo struct {
	ID          uint64 `json:"id" example:"1"`
	PhoneNumber string `json:"phone_number" example:"09121234567"`
	Name        string `json:"name" example:"Ali"`
}
type RegisterRequest struct {
	Name        string `json:"name" example:"Ali"`
	PhoneNumber string `json:"phone_number" example:"09121234567"`
	Password    string `json:"password" example:"12345678"`
}

type RegisterResponse struct {
	User UserInfo `json:"user"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" example:"09121234567"`
	Password    string `json:"password" example:"12345678"`
}

type LoginResponse struct {
	User   UserInfo `json:"user"`
	Tokens Tokens   `json:"tokens"`
}

type ProfileResponse struct {
	UserID string `json:"user_id" example:"1"`
	Name   string `json:"name" example:"Ali"`
}
type ProfileRequest struct {
	UserId string `json:"user_id"`
}
type Tokens struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIs..."`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIs..."`
}
