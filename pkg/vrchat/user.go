package vrchat

import (
	"strconv"
	"strings"

	"vrcgo/pkg/vrchat/structs"
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

func (v *Client) GetUserByID(id string) (info structs.User, err error) {
	_, err = v.client.R().
		SetResult(&info).
		Get(strings.Join([]string{"users", id}, "/"))
	return
}

func (v *Client) GetUserByName(name string) (info structs.User, err error) {
	_, err = v.client.R().
		SetResult(&info).
		Get(strings.Join([]string{"users", name, "name"}, "/"))
	return
}

func (v *Client) Search(search string, offset, n int) (list []structs.User, err error) {
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

func (v *Client) Friends(offline bool, offset, n int) (list []structs.User, err error) {
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
