package wss

import (
	"vrcgo/pkg/vrchat/structs"
)

type (
	FriendActiveEvent   func(structs.WebSocketFriendActive)
	FriendAddEvent      func(structs.WebSocketFriendAdd)
	FriendDeleteEvent   func(structs.WebSocketFriendDelete)
	FriendLocationEvent func(structs.WebSocketFriendLocation)
	FriendOfflineEvent  func(structs.WebSocketFriendOffline)
	FriendOnlineEvent   func(structs.WebSocketFriendOnline)
	FriendUpdateEvent   func(structs.WebSocketFriendUpdate)
	UserUpdateEvent     func(structs.WebSocketUserUpdate)
)

var (
	friendActiveEvent   FriendActiveEvent
	friendAddEvent      FriendAddEvent
	friendDeleteEvent   FriendDeleteEvent
	friendLocationEvent FriendLocationEvent
	friendOfflineEvent  FriendOfflineEvent
	friendOnlineEvent   FriendOnlineEvent
	friendUpdateEvent   FriendUpdateEvent
	userUpdateEvent     UserUpdateEvent
)

func (c *Client) OnFriendAdd(f FriendAddEvent) {
	friendAddEvent = f
}

func (c *Client) OnFriendDelete(f FriendDeleteEvent) {
	friendDeleteEvent = f
}

func (c *Client) OnFriendLocation(f FriendLocationEvent) {
	friendLocationEvent = f
}

func (c *Client) OnFriendActive(f FriendActiveEvent) {
	friendActiveEvent = f
}

func (c *Client) OnFriendOffline(f FriendOfflineEvent) {
	friendOfflineEvent = f
}

func (c *Client) OnFriendOnline(f FriendOnlineEvent) {
	friendOnlineEvent = f
}

func (c *Client) OnFriendUpdate(f FriendUpdateEvent) {
	friendUpdateEvent = f
}

func (c *Client) OnUserUpdate(f UserUpdateEvent) {
	userUpdateEvent = f
}
