package main

import (
	"github.com/dustinliu/taskcommander/controller"
)

func main() {
	if err := controller.GetApplication().Run(); err != nil {
		panic(err)
	}
}
