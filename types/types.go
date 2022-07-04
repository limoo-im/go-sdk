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

type SendMessageOptions struct {
	WorkspaceID          string  `json:"-"`
	ConversationID       string  `json:"-"`
	Text                 string  `json:"text"`
	UserID               *string `json:"user_id"`
	CreateAt             *int    `json:"create_at"`
	UpdateAt             *int    `json:"update_at"`
	ThreadRootID         *string `json:"thread_root_id"`
	DirectReplyMessageID *string `json:"direct_reply_message_id"`
	PendingMessageID     *string `json:"pending_message_id"`
	Files                []Files `json:"files"`
	Props                *string `json:"props"`
	TimelineLabel        *string `json:"timeline_label"`
	Type                 *string `json:"type"`
}

type Files struct {
	Hash     string            `json:"hash"`
	Size     int               `json:"size"`
	Name     string            `json:"name"`
	MimeType string            `json:"mime_type"`
	CreateAt int               `json:"create_at"`
	Metadata map[string]string `json:"metadata"`
}
