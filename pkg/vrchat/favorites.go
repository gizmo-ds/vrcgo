package vrchat

import (
	"strings"

	"vrcgo/pkg/vrchat/structs"

	"github.com/go-resty/resty/v2"
)

const (
	FavoriteTypeWorld  = "world"
	FavoriteTypeFriend = "friend"
	FavoriteTypeAvatar = "avatar"

	FavoriteFriendTag1 = "group_0"
	FavoriteFriendTag2 = "group_1"
	FavoriteFriendTag3 = "group_2"

	FavoriteWorldTag1 = "worlds1"
	FavoriteWorldTag2 = "worlds2"
	FavoriteWorldTag3 = "worlds3"
	FavoriteWorldTag4 = "worlds4"
)

type favorite struct {
	client *resty.Client
}

func (v *favorite) Get(favoriteID string) (info structs.Favorite, err error) {
	_, err = v.client.R().
		SetResult(&info).
		Get(strings.Join([]string{"favorites", favoriteID}, "/"))
	return
}

func (v *favorite) All(favoriteType string) (list []structs.Favorite, err error) {
	_, err = v.client.R().
		SetResult(&list).
		SetQueryParam("type", favoriteType).
		Get("favorites")
	return
}

func (v *favorite) Friend(tag ...string) (list []structs.LimitedUser, err error) {
	req := v.client.R().
		SetResult(&list)
	if len(tag) > 0 {
		req = req.SetQueryParam("tag", tag[0])
	}
	_, err = req.Get("auth/user/friends/favorite")
	return
}

func (v *favorite) World() (list []structs.LimitedWorld, err error) {
	_, err = v.client.R().
		SetResult(&list).
		Get("worlds/favorites")
	return
}

func (v *favorite) Add(favoriteType, objectID, tag string) (info structs.Favorite, err error) {
	_, err = v.client.R().
		SetResult(&info).
		SetBody(map[string]interface{}{
			"type":       favoriteType,
			"favoriteId": objectID,
			"tags":       []string{tag},
		}).
		Post("favorites")
	return
}

func (v *favorite) Remove(favoriteID string) (err error) {
	_, err = v.client.R().
		Delete(strings.Join([]string{"favorites", favoriteID}, "/"))
	return
}
