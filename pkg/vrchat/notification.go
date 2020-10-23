package vrchat

import (
	"strings"

	"vrcgo/pkg/vrchat/structs"

	"github.com/go-resty/resty/v2"
)

const (
	NotificationTypeRequestInvite = "requestInvite"
	NotificationTypeInvite        = "invite"

	// VRChat API not implemented
	NotificationTypeBroadcast = "broadcast"

	// VRChat API not implemented
	NotificationTypeMessage = "message"
)

type notification struct {
	client *resty.Client
}

func (v *notification) All() (list []structs.Notification, err error) {
	_, err = v.client.R().
		SetResult(&list).
		Get("auth/user/notifications")
	return
}

func (v *notification) Hide(id string) error {
	_, err := v.client.R().
		Put(strings.Join([]string{"auth/user/notifications", id, "hide"}, "/"))
	return err
}

func (v *notification) See(id string) error {
	_, err := v.client.R().
		Put(strings.Join([]string{"auth/user/notifications", id, "see"}, "/"))
	return err
}

func (v *notification) Send(userID, t string, details interface{}) error {
	_, err := v.client.R().
		SetBody(
			map[string]interface{}{
				"type":    t,
				"details": details,
			},
		).
		Post(strings.Join([]string{"user", userID, "notification"}, "/"))
	return err
}
