package vrchat

import (
	"errors"
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

func (v *Client) UserSearch(search string, offset, n int) (list []structs.User, err error) {
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

func (v *Client) Logout() error {
	_, err := v.client.R().Put("logout")
	return err
}

func (v *Client) Verify2FA(code string) error {
	var info struct {
		Verified bool `json:"verified"`
	}
	_, err := v.client.R().
		SetBody(map[string]string{"code": code}).
		SetResult(&info).
		Post("auth/twofactorauth/totp/verify")
	if err != nil {
		return err
	} else if !info.Verified {
		return errors.New("2FA verification failed")
	}
	return nil
}

func (v *Client) AuthUser() (err error) {
	_, err = v.client.R().
		SetResult(&v.User).
		Get("auth/user")
	if err != nil {
		return
	}
	return
}

func (v *Client) AuthToken() (token string, err error) {
	var info struct {
		Token string `json:"token"`
	}
	_, err = v.client.R().
		SetResult(&info).
		Get("auth")
	if err != nil {
		return
	}
	token = info.Token
	return
}
