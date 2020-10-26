package vrchat

import "github.com/go-resty/resty/v2"

type moderation struct {
	client *resty.Client
}

const (
	ModerationTypeMute    = "mute"
	ModerationTypeUnMute  = "unmute"
	ModerationTypeBlock   = "block"
	ModerationTypeUnBlock = "unblock"
)

// Returns a list of moderations you sent.
func (v *moderation) Players() {
	v.client.R().
		Get("auth/user/playermoderations")
}

// Returns a list with all the moderations against the current user.
func (v *moderation) Against() {
	v.client.R().
		Get("auth/user/playermoderated")
}
