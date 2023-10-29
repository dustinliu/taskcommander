package service

import (
	"time"

	"github.com/google/uuid"
)

type TaskType uint8

const (
	INBOX TaskType = iota
	NEXT
	SOMEDAY
	FOCUS
)

type Task struct {
	Id        string
	Title     string
	Note      string
	Type      TaskType
	Due       time.Time
	Project   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TaskService interface {
	AddTask(task *Task) (*Task, error)
	UpdateTask(task *Task) error
	GetTask(id string) (*Task, error)
	DeleteTask(id string) error
}

func getUUID() string {
	return uuid.New().String()
}
