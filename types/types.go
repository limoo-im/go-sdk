package types

type LoginResult struct {
	UID         string `json:"uid"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Id          string `json:"id"`
	Locale      string `json:"locale"`
	Position    string `json:"position"`
	Nickname    string `json:"nickname"`
	PhoneNumber string `json:"phone_number"`
	AvatarHash  string `json:"avatar_hash"`
	// CreatedAt            string   `json:"created_at"`
	// UpdateAt             string   `json:"update_at"`
	// DeleteAt             string   `json:"delete_at"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	IsBot         bool   `json:"is_bot"`
	// LastAvatarUpdateTime string   `json:"last_avatar_update_time"`
	Roles []string `json:"roles"`
}
