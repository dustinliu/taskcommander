package service

import (
	"fmt"
)

type Task interface {
	GetId() string
	GetTitle() string
	SetTitle(string) Task
	GetNote() string
	SetNote(string) Task
	GetFocus() bool
	SetFocus(bool) Task
	GetStatus() Status
	SetStatus(Status) Task
	GetProject() string
	SetProject(string) Task
	GetTags() []string
	SetTag(string) Task
	GetCategory() Category
	SetCategory(Category) Task
	GetDue() string // RFC3339
	SetDue(string) Task
	GetCompleted() string // RFC3339
	SetCompleted(string) Task
	GetUpdated() string // RFC3339
	Error() error
	String() string
}

func taskToString(t Task) string {
	return fmt.Sprintf("%T: {Id:%s Title:%s Note:%s Status:%v Project:%s Focus:%t Tags:%v Category:%s Due:%s Completed:%s Updated:%s}",
		t, t.GetId(), t.GetTitle(), t.GetNote(), t.GetStatus().Name(), t.GetProject(), t.GetFocus(), t.GetTags(), t.GetCategory().Name(), t.GetDue(), t.GetCompleted(), t.GetUpdated())
}

type TaskService interface {
	OAuth2(urlHandler func(string) error) error
	NewTask() Task
	AddTask(Task) (Task, error)
	ListTodoTasks() ([]Task, error)
	ListTags() ([]string, error)
	ListProjects() ([]string, error)
}

func NewService() (TaskService, error) {
	return NewGoogleTaskService()
}
