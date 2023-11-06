package main

import (
	"os"

	"github.com/dustinliu/taskcommander/controller"
	"github.com/dustinliu/taskcommander/service"
)

func main() {
	if err := controller.GetApplication().Run(); err != nil {
		service.GetLogger().Error(err.Error())
		os.Exit(1)
	}
}
