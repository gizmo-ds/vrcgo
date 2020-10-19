package vrchat

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"time"

	"vrcgo/pkg/vrchat/structs"
	"vrcgo/pkg/vrchat/wss"

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
		WSS    *wss.Client
		User   structs.CurrentUser
	}

	VRChat struct {
		client *resty.Client
	}
)

func New() *VRChat {
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

	return &VRChat{
		client: client,
	}
}

func (v *VRChat) Login(username, password string) (*Client, error) {
	client := *v.client
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client.SetCookieJar(jar)

	var info structs.CurrentUser
	_, err = client.R().
		SetBasicAuth(username, password).
		SetResult(&info).
		Get("auth/user")
	if err != nil {
		return nil, err
	}

	return newClient(&client, info), nil
}

func (v *VRChat) LoginWithSteam(steamTicket string) (*Client, error) {
	client := *v.client
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client.SetCookieJar(jar)

	var info structs.CurrentUser
	_, err = client.R().
		SetBody(map[string]string{"steamTicket": steamTicket}).
		SetResult(&info).
		Post("auth/steam")
	if err != nil {
		return nil, err
	}

	return newClient(&client, info), nil
}

func newClient(client *resty.Client, user structs.CurrentUser) *Client {
	return &Client{
		client: client,
		WSS:    wss.NewClient(),
		User:   user,
	}
}
