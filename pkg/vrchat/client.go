package vrchat

import (
	"errors"

	"vrcgo/pkg/vrchat/structs"
	"vrcgo/pkg/vrchat/wss"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client       *resty.Client
	WSS          *wss.Client
	CurrentUser  structs.CurrentUser
	User         *user
	Favorite     *favorite
	Notification *notification
}

func newClient(client *resty.Client, u structs.CurrentUser) *Client {
	return &Client{
		client:       client,
		WSS:          wss.NewClient(),
		CurrentUser:  u,
		User:         &user{client: client},
		Favorite:     &favorite{client: client},
		Notification: &notification{client: client},
	}
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
		SetResult(&v.CurrentUser).
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
