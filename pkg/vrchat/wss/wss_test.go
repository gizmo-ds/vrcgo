package wss_test

import (
	"os"
	"testing"

	"vrcgo/pkg/vrchat/wss"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewClientAndClose(t *testing.T) {
	_ = godotenv.Load("../.env")

	authToken := os.Getenv("VRC_AUTH_TOKEN")
	if authToken == "" {
		t.Skip("skipping test; $VRC_AUTH_TOKEN not set")
	}

	client := wss.NewClient()
	_, cancel, err := client.Dial(authToken)
	assert.NoError(t, err)
	assert.NotNil(t, cancel)
	cancel()
}
