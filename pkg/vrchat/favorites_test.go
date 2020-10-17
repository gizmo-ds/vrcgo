package vrchat_test

import (
	"testing"

	"vrcgo/pkg/vrchat"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const (
	TestFavoriteID      = "fvrt_54064522-837d-4bef-ac2d-33580faddb70"
	TestFavoriteWorldID = "wrld_91805f98-fd56-497a-bf7e-772e434efb84"
)

func TestGetFavorite(t *testing.T) {
	_ = godotenv.Load()

	client := login(t)
	info, err := client.GetFavorite(TestFavoriteID)
	assert.NoError(t, err)
	assert.Equal(t, info.FavoriteId, TestFavoriteWorldID)
}

func TestListAllFavorites(t *testing.T) {
	_ = godotenv.Load()

	client := login(t)
	list, err := client.ListAllFavorites(vrchat.FavoriteTypeWorld)
	assert.NoError(t, err)
	assert.NotNil(t, list)
}

func TestListFriendFavorites(t *testing.T) {
	_ = godotenv.Load()

	client := login(t)
	list, err := client.ListFriendFavorites(vrchat.FavoriteFriendTag1)
	assert.NoError(t, err)
	assert.NotNil(t, list)
}

func TestListWorldFavorites(t *testing.T) {
	_ = godotenv.Load()

	client := login(t)
	list, err := client.ListWorldFavorites()
	assert.NoError(t, err)
	assert.NotNil(t, list)
}
