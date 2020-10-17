package vrchat_test

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const (
	TestUserID   = "usr_3c5b20e5-ee00-4b77-8943-e0ef213d689a"
	TestUserName = "gizmo_"
)

func TestGetUserByID(t *testing.T) {
	_ = godotenv.Load()

	user := login(t)
	userInfo, err := user.GetUserByID(TestUserID)
	assert.NoError(t, err)
	assert.Equal(t, userInfo.Username, TestUserName)
}

func TestGetUserByName(t *testing.T) {
	_ = godotenv.Load()

	user := login(t)
	userInfo, err := user.GetUserByName(TestUserName)
	assert.NoError(t, err)
	assert.Equal(t, userInfo.Username, TestUserName)
}
