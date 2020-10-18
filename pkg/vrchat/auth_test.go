package vrchat_test

import (
	"os"
	"testing"

	"vrcgo/pkg/vrchat"
	"vrcgo/pkg/vrchat/structs"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func login(t *testing.T) *vrchat.Client {
	client := vrchat.New()
	user, err := client.Login(os.Getenv("VRC_USN"), os.Getenv("VRC_PWD"))
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.Auth.Username, os.Getenv("VRC_USN"))
	return user
}

func login2FA(t *testing.T) *vrchat.Client {
	client := vrchat.New()
	user, err := client.Login(os.Getenv("VRC_USN_2FA"), os.Getenv("VRC_PWD_2FA"))
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, user.Auth.RequiresTwoFactorAuth)
	return user
}

func TestLogout(t *testing.T) {
	_ = godotenv.Load()

	user := login(t)
	err := user.Logout()
	assert.NoError(t, err)
}

func TestVerify2FA(t *testing.T) {
	_ = godotenv.Load()

	code := os.Getenv("VRC_2FA_CODE")
	if code == "" {
		t.Skip("skipping test; $VRC_2FA_CODE not set")
	}

	user := login2FA(t)
	assert.NotNil(t, user.Auth.RequiresTwoFactorAuth)
	err := user.Verify2FA(code)
	assert.NoError(t, err)
}

func TestAuthUser(t *testing.T) {
	_ = godotenv.Load()

	user := login(t)
	user.Auth = structs.CurrentUser{}
	err := user.AuthUser()
	assert.NoError(t, err)
	assert.Equal(t, user.Auth.Username, os.Getenv("VRC_USN"))
}
