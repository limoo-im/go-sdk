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
	Hash            *string     `json:"hash"`
	Size            *int        `json:"size"`
	Name            *string     `json:"name"`
	MimeType        *string     `json:"mime_type"`
	Extension       *string     `json:"extension"`
	HasPreviewImage *bool       `json:"has_preview_image"`
	Metadata        map[any]any `json:"metadata"`
}

type SendMessageResponse struct {
	ID                   *string               `json:"id"`
	CreateAt             *int                  `json:"create_at"`
	DeleteAt             *int                  `json:"delete_at"`
	EditAt               *int                  `json:"edit_at"`
	UserID               *string               `json:"user_id"`
	ConversationID       *string               `json:"conversation_id"`
	ThreadRootID         *string               `json:"thread_root_id"`
	Text                 *string               `json:"text"`
	Type                 *string               `json:"type"`
	Props                map[any]any           `json:"props"`
	Files                []Files               `json:"files"`
	ForwardSourceMessage *ForwardSourceMessage `json:"forward_source_message"`
	DirectReplyMessage   *DirectReplyMessage   `json:"direct_reply_message"`
	DirectReplyMessageID *string               `json:"direct_reply_message_id"`
	PendingMessageID     *string               `json:"pending_message_id"`
	Thread               *Thread               `json:"thread"`
	Reactions            []Reactions           `json:"reactions"`
	TimelineLabel        *string               `json:"timeline_label"`
	WorkspaceID          *string               `json:"workspace_id"`
	WorkspaceDisplayName *string               `json:"workspace_display_name"`
}

type ForwardSourceMessage struct {
	ID       *string `json:"id"`
	SenderID *string `json:"sender_id"`
}
type DirectReplyMessage struct {
	ID       *string `json:"id"`
	Text     *string `json:"text"`
	UserID   *string `json:"user_id"`
	DeleteAt *int    `json:"delete_at"`
	Type     *string `json:"type"`
}
type Participants struct {
	ID  *string `json:"id"`
	UID *string `json:"uid"`
}

type Thread struct {
	UID              *string        `json:"uid"`
	RootMessageID    *string        `json:"root_message_id"`
	ConversationID   *string        `json:"conversation_id"`
	MessageCount     *int           `json:"message_count"`
	Participants     []Participants `json:"participants"`
	LastActivityDate *int           `json:"last_activity_date"`
	IsUnread         *bool          `json:"is_unread"`
	IsMentioned      *bool          `json:"is_mentioned"`
	IsFollowed       *bool          `json:"is_followed"`
	Messages         map[any]any    `json:"messages"`
}
type Reactions struct {
	UserID    *string `json:"user_id"`
	EmojiName *string `json:"emoji_name"`
}
