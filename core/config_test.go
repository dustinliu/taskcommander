package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/adrg/xdg"
	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	os.Setenv("TC_RUNTIME_ENV", "dev")
	defer os.Unsetenv("TC_RUNTIME_ENV")

	config := GetConfig()
	assert.Equal(t, EnvDev, config.Env)
	assert.Equal(t, "gtask", config.Backend)
	assert.Equal(t, filepath.Join(xdg.ConfigHome, AppName, "credential.json"), config.Gtask.CredentialFile)
}
