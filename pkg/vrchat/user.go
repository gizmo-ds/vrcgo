package vrchat

import (
	"strconv"
	"strings"

	"vrcgo/pkg/vrchat/structs"

	"github.com/go-resty/resty/v2"
)

const (
	UserStateOnline  = "online"
	UserStateActive  = "active"
	UserStateOffline = "offline"

	UserStatusActive  = "active"
	UserStatusJoinMe  = "join me"
	UserStatusAskMe   = "ask me"
	UserStatusBusy    = "busy"
	UserStatusOffline = "offline"
)

type user struct {
	client *resty.Client
}

func (v *user) GetByID(id string) (info structs.User, err error) {
	_, err = v.client.R().
		SetResult(&info).
		Get(strings.Join([]string{"users", id}, "/"))
	return
}

func (v *user) GetByName(name string) (info structs.User, err error) {
	_, err = v.client.R().
		SetResult(&info).
		Get(strings.Join([]string{"users", name, "name"}, "/"))
	return
}

func (v *user) Search(search string, offset, n int) (list []structs.User, err error) {
	_, err = v.client.R().
		SetQueryParams(map[string]string{
			"offset": strconv.Itoa(offset),
			"n":      strconv.Itoa(n),
			"search": search,
		}).
		SetResult(&list).
		Get("users")
	return
}

func (v *user) Friends(offline bool, offset, n int) (list []structs.User, err error) {
	_, err = v.client.R().
		SetQueryParams(map[string]string{
			"offset": strconv.Itoa(offset),
			"n":      strconv.Itoa(n),
			"offline": func() string {
				if offline {
					return "true"
				}
				return "false"
			}(),
		}).
		SetResult(&list).
		Get("auth/user/friends")
	return
}
