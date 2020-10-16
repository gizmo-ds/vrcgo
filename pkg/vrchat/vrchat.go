package vrchat

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"time"
	"vrcgo/pkg/vrchat/structs"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

const (
	ClientUserAgent = "vrcgo/version 0.1.0"
	VRChatAPI       = "https://api.vrchat.cloud/api/1"
)

type (
	Client struct {
		client *resty.Client
		Auth   structs.AuthResponse
	}

	VRChat struct {
		client *resty.Client
	}
)

func New() (*VRChat, error) {
	client := resty.New()

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		if resp.StatusCode() != http.StatusOK {
			return fmt.Errorf(
				"Response Error: %s\n%s\n",
				resp.Status(), resp.String(),
			)
		}
		return nil
	})

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	client.JSONMarshal = json.Marshal
	client.JSONUnmarshal = json.Unmarshal

	client.
		SetHostURL(VRChatAPI).
		SetHeader("User-Agent", ClientUserAgent).
		SetTimeout(time.Second * 10).
		SetRetryCount(3)

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client.SetCookieJar(jar)

	v := VRChat{
		client: client,
	}

	err = v.RemoteConfig()

	return &v, err
}

func (v *VRChat) RemoteConfig() error {
	_, err := v.client.R().Get("config")
	return err
}

func (v *VRChat) Login(username, password string) (user *Client, err error) {
	var info structs.AuthResponse
	_, err = v.client.R().
		SetBasicAuth(username, password).
		SetResult(&info).
		Get("auth/user")
	if err != nil {
		return nil, err
	}
	user = &Client{
		client: v.client,
		Auth:   info,
	}
	return
}

func (v *VRChat) LoginWithSteam(steamTicket string) (user *Client, err error) {
	var info structs.AuthResponse
	_, err = v.client.R().
		SetBody(map[string]string{"steamTicket": steamTicket}).
		SetResult(&info).
		Post("auth/steam")
	if err != nil {
		return nil, err
	}
	user = &Client{
		client: v.client,
		Auth:   info,
	}
	return
}
