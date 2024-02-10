package core

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/adrg/xdg"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type Env uint8

const (
	EnvKey = "TC_RUNTIME_ENV"

	EnvDev Env = iota
	EnvProd

	backendKey             = "backend"
	gtaskCredentialFileKey = "gtask.credential_file"
	DebugKey               = "debug"
)

var (
	configOnce sync.Once
	config     Config
)

type gtaskConfig struct {
	CredentialFile string
}

type Config struct {
	Env     Env
	Debug   bool
	Backend string
	Gtask   gtaskConfig
}

// TODO: remove panic
func InitConfig() {
	configOnce.Do(func() {
		switch os.Getenv(EnvKey) {
		case "dev":
			config.Env = EnvDev
		default:
			config.Env = EnvProd
		}

		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		if config.Env == EnvDev {
			root, err := GetProjectRoot()
			if err != nil {
				panic(fmt.Errorf("failed to get project root: %w", err))
			}
			viper.AddConfigPath(filepath.Join(root, "tests"))
		} else {
			viper.AddConfigPath(filepath.Join(xdg.ConfigHome, AppName))
		}
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}

		config.Backend = viper.GetString(backendKey)

		f, err := homedir.Expand(viper.GetString(gtaskCredentialFileKey))
		if err != nil {
			panic(fmt.Errorf("failed to expand gtask.credential_file: %w", err))
		}
		config.Gtask.CredentialFile = f

		config.Debug = viper.GetBool(DebugKey)
	})
}

func GetConfig() Config {
	InitConfig()
	return config
}

func GetProjectRoot() (string, error) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	root, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	return root, nil
}
