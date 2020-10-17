package vrchat_test

import (
	"os"
	"testing"
	"vrcgo/pkg/vrchat"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	godotenv.Load()

	client := vrchat.New()
	user, err := client.Login(os.Getenv("VRC_USN"), os.Getenv("VRC_PWD"))
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.Auth.Username, os.Getenv("VRC_USN"))
}

func TestLoginWithSteam(t *testing.T) {
	godotenv.Load()

	steamTicket := os.Getenv("STEAM_TICKET")
	if steamTicket == "" {
		t.Skip("skipping test; $STEAM_TICKET not set")
	}

	client := vrchat.New()
	user, err := client.LoginWithSteam(steamTicket)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.Auth.Username, os.Getenv("VRC_USN"))
}
