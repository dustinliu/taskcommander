package gtask

import (
	"time"

	"github.com/dustinliu/taskcommander/service"
	tasks "google.golang.org/api/tasks/v1"
)

type GoogleTask struct {
	*tasks.Task
	focus    bool
	project  string
	tags     []string
	category service.Category
}

func NewGoogleTask() *GoogleTask {
	return &GoogleTask{
		Task:     &tasks.Task{},
		tags:     []string{},
		category: service.CategoryInbox,
	}
}

func (t *GoogleTask) GetId() string {
	return t.Id
}

func (t *GoogleTask) GetTitle() string {
	return t.Title
}

func (t *GoogleTask) SetTitle(title string) service.Task {
	t.Title = title
	return t
}

func (t *GoogleTask) GetNotes() string {
	return t.Notes
}

func (t *GoogleTask) SetNotes(notes string) service.Task {
	t.Notes = notes
	return t
}

func (t *GoogleTask) GetFocus() bool {
	return t.focus
}

func (t *GoogleTask) SetFocus(focus bool) service.Task {
	t.focus = focus
	return t
}

func (t *GoogleTask) GetStatus() service.Status {
	if t.Status == "completed" {
		return service.StatusDone
	} else {
		return service.StatusTodo
	}
}

func (t *GoogleTask) SetStatus(status service.Status) service.Task {
	if status == service.StatusDone {
		t.Status = "completed"
	} else {
		t.Status = "needsAction"
	}

	return t
}

func (t *GoogleTask) GetProject() string {
	return t.project
}

func (t *GoogleTask) SetProject(project string) service.Task {
	t.project = project
	return t
}

func (t *GoogleTask) GetTags() []string {
	return t.tags
}

func (t *GoogleTask) SetTag(tag string) service.Task {
	t.tags = append(t.tags, tag)
	return t
}

func (t *GoogleTask) GetCategory() service.Category {
	return t.category
}

func (t *GoogleTask) SetCategory(category service.Category) service.Task {
	t.category = category
	return t
}

func (t *GoogleTask) GetDue() time.Time {
	due, err := time.Parse(time.RFC3339, t.Due)
	if err != nil {
		return time.Time{}
	}
	return due
}

func (t *GoogleTask) SetDue(due time.Time) service.Task {
	t.Due = due.Format(time.RFC3339)
	return t
}

func (t *GoogleTask) GetCompleted() time.Time {
	if t.Completed == nil {
		return time.Time{}
	}

	c, err := time.Parse(time.RFC3339, *t.Completed)
	if err != nil {
		return time.Time{}
	}
	return c
}

func (t *GoogleTask) SetCompleted(comp time.Time) service.Task {
	c := comp.Format(time.RFC3339)
	t.Completed = &c
	return t
}
