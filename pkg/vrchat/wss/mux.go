package wss

import (
	"encoding/json"
	"log"

	"vrcgo/pkg/vrchat/structs"
)

func mux(msg structs.WebSocketMessage) {
	var unmarshal = func(info interface{}) {
		if err := json.Unmarshal([]byte(msg.Content), &info); err != nil {
			log.Println(err)
			return
		}
	}

	switch msg.Type {
	case "notification":
		var info structs.Notification
		unmarshal(&info)
		notification(info)

	case "friend-active":
		var info structs.WebSocketFriendActive
		unmarshal(&info)
		if friendActiveEvent != nil {
			friendActiveEvent(info)
		}

	case "friend-add":
		var info structs.WebSocketFriendAdd
		unmarshal(&info)
		if friendAddEvent != nil {
			friendAddEvent(info)
		}

	case "friend-delete":
		var info structs.WebSocketFriendDelete
		unmarshal(&info)
		if friendDeleteEvent != nil {
			friendDeleteEvent(info)
		}

	case "friend-location":
		var info structs.WebSocketFriendLocation
		unmarshal(&info)
		if friendLocationEvent != nil {
			friendLocationEvent(info)
		}

	case "friend-offline":
		var info structs.WebSocketFriendOffline
		unmarshal(&info)
		if friendOfflineEvent != nil {
			friendOfflineEvent(info)
		}

	case "friend-online":
		var info structs.WebSocketFriendOnline
		unmarshal(&info)
		if friendOnlineEvent != nil {
			friendOnlineEvent(info)
		}

	case "friend-update":
		var info structs.WebSocketFriendUpdate
		unmarshal(&info)
		if friendUpdateEvent != nil {
			friendUpdateEvent(info)
		}

	case "user-update":
		var info structs.WebSocketUserUpdate
		unmarshal(&info)
		if userUpdateEvent != nil {
			userUpdateEvent(info)
		}

	default:
		log.Printf("Default: %v\n%v", msg.Type, msg.Content)
	}
}

func notification(info structs.Notification) {
	switch info.Type {
	case "invite":
		log.Printf("Invite: %s -> %s\n", info.SenderUsername, info.Invite().WorldName)

	case "friendRequest":
		log.Printf("FriendRequest: %s %s\n", info.SenderUsername, info.RequestInvite().Platform)

	default:
		log.Printf("Notification: %v\n%+v", info.Type, info)
	}
}
