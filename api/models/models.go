package models

type NOT_IMPLEMENTED string

type NewUser struct {
	Email       *string `json:"email,omitempty"`
	Credentials *string `json:"credentials,omitempty"`
}

type User struct {
	UserID *string `json:"user_id,omitempty"`
	Email  *string `json:"email,omitempty"`
}

type AuthenticatedUser struct {
	User   *User   `json:"user,omitempty"`
	Token  *string `json:"token,omitempty"`
	Secret *string `json:"secret,omitempty"`
}

type ChannelGroup struct {
	GroupID    *string  `json:"group_id,omitempty"`
	UserID     *string  `json:"user_id,omitempty"`
	GroupName  *string  `json:"group_name,omitempty"`
	GroupDescr *string  `json:"group_descr,omitempty"`
	ChannelIDs []int    `json:"channel_ids,omitempty"`
	Tags       []string `json:"tags,omitempty"`
	Order      *int     `json:"order,omitempty"`
	CategoryID *string  `json:"category_id,omitempty"`
}

type GroupCategory struct {
	CategoryID    *string `json:"category_id,omitempty"`
	UserID        *string `json:"user_id,omitempty"`
	CategoryName  *string `json:"category_name,omitempty"`
	CategoryDescr *string `json:"category_descr,omitempty"`
	Order         *int    `json:"order,omitempty"`
	Favorite      *bool   `json:"favorite,omitempty"`
}

type TelegramChannel struct {
	ID          *int             `json:"id,omitempty"`
	Title       *string          `json:"title,omitempty"`
	Photo       *string          `json:"photo,omitempty"`
	FromSession *TelegramSession `json:"from_session,omitempty"`
}

type TelegramSession struct {
	Phone *string `json:"phone,omitempty"`
}

type TelegramMessage struct {
	ID       *int             `json:"id,omitempty"`
	Channel  *TelegramChannel `json:"channel,omitempty"`
	FromName *string          `json:"from_name,omitempty"`
	Date     *int             `json:"date,omitempty"`
	Message  *string          `json:"message,omitempty"`
	Attached *NOT_IMPLEMENTED `json:"attached,omitempty"`
	TTL      *int             `json:"ttl_period,omitempty"`
}

type TelegramConfirmationCode struct {
	PhoneCode *string `json:"phone_code,omitempty"`
	Key       *string `json:"key,omitempty"`
}

type LinkStatus struct {
	Ok bool
	PasswordRequired bool
}