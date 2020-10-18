package vrchat

import (
	"errors"
)

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
		SetResult(&v.Auth).
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
