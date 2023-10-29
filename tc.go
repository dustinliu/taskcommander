package main

import "github.com/dustinliu/taskcommander/view"

func main() {
	if err := view.GetApplication().Run(); err != nil {
		panic(err)
	}
}
