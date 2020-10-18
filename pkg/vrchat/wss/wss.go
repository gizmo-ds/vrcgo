package wss

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"vrcgo/pkg/vrchat/structs"

	"github.com/gorilla/websocket"
)

const (
	VRChatWebSocketURL = "wss://pipeline.vrchat.cloud/"
)

type (
	Client struct {
		dialer *websocket.Dialer
		conn   *websocket.Conn
		ctx    context.Context
		cancel func()
	}
)

func NewClient() *Client {
	return &Client{
		dialer: &websocket.Dialer{
			Proxy:            http.ProxyFromEnvironment,
			HandshakeTimeout: 45 * time.Second,
		},
	}
}

func (c *Client) Dial(authToken string) (ctx context.Context, cancel func(), err error) {
	if c.conn != nil {
		_ = c.conn.Close()
	}
	values := url.Values{}
	values.Add("authToken", authToken)
	_url, _ := url.Parse(VRChatWebSocketURL)
	_url.RawQuery = values.Encode()
	conn, _, err := c.dialer.Dial(_url.String(), nil)
	if err != nil {
		return
	}
	c.conn = conn
	ctx, cancel = context.WithCancel(context.Background())
	return
}

func (c *Client) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				_ = c.conn.Close()
				return

			default:
				msgType, msg, err := c.conn.ReadMessage()
				if err != nil {
					log.Println(err)
					if c.conn != nil {
						_ = c.conn.Close()
					}
					return
				}
				if msg != nil {
					switch msgType {
					case websocket.TextMessage:
						var info structs.WebSocketMessage
						if err := json.Unmarshal(msg, &info); err != nil {
							log.Println(err)
							return
						}
						mux(info)
					}
				}
			}
		}
	}()
}
