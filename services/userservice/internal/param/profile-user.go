package param

type ProfileRequest struct {
	UserID uint `json:"user_id"`
}

type ProfileResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
