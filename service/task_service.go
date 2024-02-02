package service

import (
	"time"
)

type (
	Status   uint8
	Category int8
)

const (
	StatusTodo Status = iota
	StatusDone

	CategoryNone  Category = -1
	CategoryInbox Category = iota
	CategoryNext
	CategorySomeday
	CategoryFocus
)

var categoryNames = []string{
	"Inbox",
	"Next",
	"Someday",
	"Focus",
}

func (c Category) IsValid() bool {
	return c >= 0 && int(c) < len(categoryNames)
}

func (c Category) Name() string {
	if !c.IsValid() {
		return "Invalid"
	}

	return categoryNames[c]
}

type Task interface {
	GetId() string
	GetTitle() string
	SetTitle(title string) Task
	GetNotes() string
	SetNotes(notes string) Task
	GetFocus() bool
	SetFocus(focus bool) Task
	GetStatus() Status
	SetStatus(status Status) Task
	GetProject() string
	SetProject(project string) Task
	GetTags() []string
	SetTag(tag string) Task
	GetCategory() Category
	SetCategory(category Category) Task
	GetDue() time.Time
	SetDue(due time.Time) Task
	GetCompleted() time.Time
	SetCompleted(time.Time) Task
}

func NewTask() Task {
	return NewGoogleTask()
}

type TaskService interface {
	InitOauth2Needed() bool
	GetOauthAuthUrl() string
	WaitForAuthDone() error
	Init() error
	AddTask(task Task) (Task, error)
	ListTasksByCategory(cat Category) ([]Task, error)
	ListTags() ([]string, error)
	ListProjects() ([]string, error)
}

func NewService() (TaskService, error) {
	s, err := NewGoogleTaskService()
	if err != nil {
		return nil, err
	}
	return s, nil
}
