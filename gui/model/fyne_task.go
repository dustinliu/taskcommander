package model

import (
	"fyne.io/fyne/v2/data/binding"
	"github.com/dustinliu/taskcommander/service"
)

var _ service.Task = (*FyneTask)(nil)

type FyneTask struct {
	title binding.String
}

func NewFyneTask(task service.Task) *FyneTask {
	t := &FyneTask{
		title: binding.NewString(),
	}
	setBindingString(task.GetTitle(), t.title)

	return t
}

// Error implements service.Task.
func (*FyneTask) Error() error {
	panic("unimplemented")
}

// GetCategory implements service.Task.
func (*FyneTask) GetCategory() service.Category {
	panic("unimplemented")
}

// GetCompleted implements service.Task.
func (*FyneTask) GetCompleted() string {
	panic("unimplemented")
}

// GetDue implements service.Task.
func (*FyneTask) GetDue() string {
	panic("unimplemented")
}

// GetFocus implements service.Task.
func (*FyneTask) GetFocus() bool {
	panic("unimplemented")
}

// GetId implements service.Task.
func (*FyneTask) GetId() string {
	panic("unimplemented")
}

// GetNote implements service.Task.
func (*FyneTask) GetNote() string {
	panic("unimplemented")
}

// GetProject implements service.Task.
func (*FyneTask) GetProject() string {
	panic("unimplemented")
}

// GetStatus implements service.Task.
func (*FyneTask) GetStatus() service.Status {
	panic("unimplemented")
}

// GetTags implements service.Task.
func (*FyneTask) GetTags() []string {
	panic("unimplemented")
}

// GetTitle implements service.Task.
func (t *FyneTask) GetTitle() string {
	return getBindingString(t.title)
}

func (t *FyneTask) GetBindingTitle() binding.String {
	return t.title
}

// GetUpdated implements service.Task.
func (*FyneTask) GetUpdated() string {
	panic("unimplemented")
}

// SetCategory implements service.Task.
func (*FyneTask) SetCategory(service.Category) service.Task {
	panic("unimplemented")
}

// SetCompleted implements service.Task.
func (*FyneTask) SetCompleted(string) service.Task {
	panic("unimplemented")
}

// SetDue implements service.Task.
func (*FyneTask) SetDue(string) service.Task {
	panic("unimplemented")
}

// SetFocus implements service.Task.
func (*FyneTask) SetFocus(bool) service.Task {
	panic("unimplemented")
}

// SetNote implements service.Task.
func (*FyneTask) SetNote(string) service.Task {
	panic("unimplemented")
}

// SetProject implements service.Task.
func (*FyneTask) SetProject(string) service.Task {
	panic("unimplemented")
}

// SetStatus implements service.Task.
func (*FyneTask) SetStatus(service.Status) service.Task {
	panic("unimplemented")
}

// SetTag implements service.Task.
func (*FyneTask) SetTag(string) service.Task {
	panic("unimplemented")
}

// SetTitle implements service.Task.
func (t *FyneTask) SetTitle(title string) service.Task {
	setBindingString(title, t.title)
	return t
}

func setBindingString(source string, target binding.String) {
	target.Set(source) // nolint:errcheck // set seems never return error
}

func getBindingString(source binding.String) string {
	s, err := source.Get()
	if err != nil {
		return err.Error()
	}
	return s
}

func (t *FyneTask) String() string {
	return t.GetTitle()
}
