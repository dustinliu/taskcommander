package service

import (
	"time"

	tasks "google.golang.org/api/tasks/v1"
)

var _ Task = (*GoogleTask)(nil)

type GoogleTask struct {
	*tasks.Task
	focus    bool
	project  string
	tags     []string
	category Category
}

func NewGoogleTask() *GoogleTask {
	return &GoogleTask{
		Task:     &tasks.Task{},
		tags:     []string{},
		category: CategoryInbox,
	}
}

func (t *GoogleTask) GetId() string {
	return t.Id
}

func (t *GoogleTask) GetTitle() string {
	return t.Title
}

func (t *GoogleTask) SetTitle(title string) Task {
	t.Title = title
	return t
}

func (t *GoogleTask) GetNotes() string {
	return t.Notes
}

func (t *GoogleTask) SetNotes(notes string) Task {
	t.Notes = notes
	return t
}

func (t *GoogleTask) GetFocus() bool {
	return t.focus
}

func (t *GoogleTask) SetFocus(focus bool) Task {
	t.focus = focus
	return t
}

func (t *GoogleTask) GetStatus() Status {
	if t.Status == "completed" {
		return StatusDone
	} else {
		return StatusTodo
	}
}

func (t *GoogleTask) SetStatus(status Status) Task {
	if status == StatusDone {
		t.Status = "completed"
	} else {
		t.Status = "needsAction"
	}

	return t
}

func (t *GoogleTask) GetProject() string {
	return t.project
}

func (t *GoogleTask) SetProject(project string) Task {
	t.project = project
	return t
}

func (t *GoogleTask) GetTags() []string {
	return t.tags
}

func (t *GoogleTask) SetTag(tag string) Task {
	t.tags = append(t.tags, tag)
	return t
}

func (t *GoogleTask) GetCategory() Category {
	return t.category
}

func (t *GoogleTask) SetCategory(category Category) Task {
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

func (t *GoogleTask) SetDue(due time.Time) Task {
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

func (t *GoogleTask) SetCompleted(comp time.Time) Task {
	c := comp.Format(time.RFC3339)
	t.Completed = &c
	return t
}
