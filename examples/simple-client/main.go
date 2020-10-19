package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"vrcgo/pkg/vrchat"
	"vrcgo/pkg/vrchat/structs"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	vrc, err := vrchat.New().Login(os.Getenv("VRC_USN"), os.Getenv("VRC_PWD"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s\n", vrc.User.DisplayName)

	token, err := vrc.AuthToken()
	if err != nil {
		panic(err)
	}

	wsClient := vrc.WSS
	cancel, err := wsClient.Start(token)
	if err != nil {
		panic(err)
	}

	wsClient.OnFriendActive(func(info structs.WebSocketFriendActive) {
		log.Printf("FriendActive: %s\n", info.User.DisplayName)
	})
	wsClient.OnFriendOffline(func(info structs.WebSocketFriendOffline) {
		log.Printf("FriendOffline: %s\n", info.UserID)
	})
	wsClient.OnFriendOnline(func(info structs.WebSocketFriendOnline) {
		log.Printf("FriendOnline: %s\n", info.User.DisplayName)
	})
	wsClient.OnFriendLocation(func(info structs.WebSocketFriendLocation) {
		log.Printf("FriendLocation: %s -> %s\n",
			info.User.DisplayName, info.World.GetName(),
		)
	})

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	<-interrupt
	cancel()
	fmt.Println("Close")
}
