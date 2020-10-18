package vrchat

import (
	"strings"
	"vrcgo/pkg/vrchat/structs"
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

func (v *Client) GetFavorite(favoriteID string) (info structs.Favorite, err error) {
	_, err = v.client.R().
		SetResult(&info).
		Get(strings.Join([]string{"favorites", favoriteID}, "/"))
	return
}

func (v *Client) ListAllFavorites(favoriteType string) (list []structs.Favorite, err error) {
	_, err = v.client.R().
		SetResult(&list).
		SetQueryParam("type", favoriteType).
		Get("favorites")
	return
}

func (v *Client) ListFriendFavorites(tag string) (list []structs.Favorite, err error) {
	_, err = v.client.R().
		SetResult(&list).
		Get("auth/user/friends/favorite")
	return
}

func (v *Client) ListWorldFavorites() (list []structs.Favorite, err error) {
	_, err = v.client.R().
		SetResult(&list).
		Get("worlds/favorites")
	return
}

func (v *Client) AddFavorite(favoriteType, objectID, tag string) (info structs.Favorite, err error) {
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

func (v *Client) RemoveFavorite(favoriteID string) (err error) {
	_, err = v.client.R().
		Delete(strings.Join([]string{"favorites", favoriteID}, "/"))
	return
}
