package controller

import (
	"github.com/dustinliu/taskcommander/gtask"
	"github.com/dustinliu/taskcommander/service"
)

func NewService() (service.TaskService, error) {
	s, err := gtask.NewGoogleTaskService()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func NewTask() service.Task {
	return gtask.NewGoogleTask()
}
