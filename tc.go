package main

import (
	"os"

	"github.com/dustinliu/taskcommander/controller"
	"github.com/dustinliu/taskcommander/logger"
)

func main() {
	if err := controller.NewApplication().Run(); err != nil {
		logger.GetLogger().Error(err.Error())
		os.Exit(1)
	}
}
