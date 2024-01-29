package controller

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
	EnvDev Env = iota
	EnvProd
)

var (
	once   sync.Once
	config Config
)

type GtaskConfig struct {
	CredentialFile string
}

type Config struct {
	Env     Env
	Backend string
	Gtask   GtaskConfig
}

func GetConfig() Config {
	once.Do(func() {
		switch os.Getenv("TC_RUNTIME_ENV") {
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
			viper.AddConfigPath(filepath.Join(root, "test"))
		} else {
			viper.AddConfigPath(filepath.Join(xdg.ConfigHome, "taskcommander"))
			viper.AddConfigPath("$HOME/.taskcommander")
		}
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}

		config.Backend = viper.GetString("backend")
		f, err := homedir.Expand(viper.GetString("gtask.credential_file"))
		if err != nil {
			panic(fmt.Errorf("failed to expand gtask.credential_file: %w", err))
		}
		config.Gtask.CredentialFile = f
	})

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
