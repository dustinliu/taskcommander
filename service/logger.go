package service

import (
	"log"
	"os"

	"github.com/adrg/xdg"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

// TODO different log for development and production?
func init() {
	env := os.Getenv("TC_RUNTIME_ENV")
	log.Print("env: ", env)
	var logPath string
	switch env {
	case "dev":
		logPath = "./tc.log"
	default:
		logPath = xdg.StateHome + "/taskcommander/tc.log"
	}

	log.Print("log path: ", logPath)
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{logPath}
	config.DisableCaller = false
	config.DisableStacktrace = false
	config.EncoderConfig.FunctionKey = "func"
	config.EncoderConfig.CallerKey = ""
	l, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}
	logger = l.Sugar()
}

func GetLogger() *zap.SugaredLogger {
	return logger
}
