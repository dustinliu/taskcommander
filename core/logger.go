package core

import (
	"log"
	"path/filepath"
	"sync"

	"github.com/adrg/xdg"
	"go.uber.org/zap"
)

var (
	Debug   = false
	logger  *zap.SugaredLogger
	logOnce sync.Once
)

// TODO: production mode
func InitLogger() {
	logFile, err := xdg.StateFile(filepath.Join(AppName, "taskcommander.log"))
	if err != nil {
		log.Fatal(err)
	}

	var config zap.Config
	if Debug {
		config = zap.NewDevelopmentConfig()
		config.DisableCaller = false
		config.DisableStacktrace = false
		config.EncoderConfig.FunctionKey = "func"
		config.EncoderConfig.CallerKey = ""
	} else {
		config = zap.NewProductionConfig()
	}
	config.OutputPaths = []string{logFile}

	l, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}
	logger = l.Sugar()
}

func GetLogger() *zap.SugaredLogger {
	logOnce.Do(func() {
		InitLogger()
	})
	return logger
}
