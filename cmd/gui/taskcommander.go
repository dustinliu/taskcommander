package main

import (
	"os"

	"github.com/dustinliu/taskcommander/core"
	"github.com/dustinliu/taskcommander/gui/controller"
)

func main() {
	app, err := controller.NewApplication()
	if err != nil {
		core.GetLogger().Error(err.Error())
		os.Exit(1)
	}
	app.Run()
}
