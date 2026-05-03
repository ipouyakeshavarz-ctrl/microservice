package param

type ProfileRequest struct {
	UserID uint
}

type ProfileResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
