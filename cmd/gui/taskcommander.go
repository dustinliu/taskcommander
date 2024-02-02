package main

import (
	"os"

	"github.com/dustinliu/taskcommander/controller"
	"github.com/dustinliu/taskcommander/core"
)

func main() {
	app, err := controller.NewApplication()
	if err != nil {
		core.GetLogger().Error(err.Error())
		os.Exit(1)
	}
	app.Run()
}
