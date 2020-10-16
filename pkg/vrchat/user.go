package vrchat

import (
	"strings"
	"vrcgo/pkg/vrchat/structs"
)

func (v *VRChatClient) GetUserByID(id string) (info structs.UserInfo, err error) {
	_, err = v.client.R().
		SetResult(&info).
		Get(strings.Join([]string{"users", id}, "/"))
	return
}

func (v *VRChatClient) GetUserByName(name string) (info structs.UserInfo, err error) {
	_, err = v.client.R().
		SetResult(&info).
		Get(strings.Join([]string{"users", name, "name"}, "/"))
	return
}
