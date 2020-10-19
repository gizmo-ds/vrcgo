package vrchat

import (
	"strings"

	"vrcgo/pkg/vrchat/structs"

	"github.com/go-resty/resty/v2"
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

func (v *notification) Delete(id string) (err error) {
	_, err = v.client.R().
		Put(strings.Join([]string{"auth/user/notifications", id, "hide"}, "/"))
	return
}
