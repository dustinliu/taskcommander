package controller

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	os.Setenv("TC_RUNTIME_ENV", "dev")
	defer os.Unsetenv("TC_RUNTIME_ENV")

	config := GetConfig()
	assert.Equal(t, EnvDev, config.Env)
	assert.Equal(t, "gtask", config.Backend)
	assert.Equal(t, "/taskcommander/google_service.json", config.Gtask.CredentialFile)
}
