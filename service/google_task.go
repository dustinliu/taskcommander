package service

import (
	"encoding/json"
	"fmt"
	"strings"

	tasks "google.golang.org/api/tasks/v1"
)

const (
	note_seperator = "\n--- TASK_SEPERATOR ---\n"

	gtaskStatusDone    = "completed"
	gtaskStatusTodo    = "needsAction"
	gtaskStatusInvalid = "invalid"
)

type internalStatus struct {
	Focus    bool     `json:"focus"`
	Project  string   `json:"project"`
	Tags     []string `json:"tags"`
	Category Category `json:"category"`
}

type googleTask struct {
	*tasks.Task
	note string
	internalStatus
	err error
}

func newGoogleTask(t *tasks.Task) *googleTask {
	task := &googleTask{Task: t}

	igtask := internalStatus{Category: CategoryInbox}
	strs := strings.Split(t.Notes, note_seperator)
	task.note = strs[0]
	if len(strs) > 1 {
		if err := json.Unmarshal([]byte(strs[1]), &igtask); err != nil {
			task.err = fmt.Errorf("failed to unmarshal internal status: %v", err)
			return task
		}
		task.internalStatus = igtask
	}

	return task
}

func (g *googleTask) getGoogleTask() (*tasks.Task, error) {
	statusJson, err := json.Marshal(g.internalStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal internal status: %v", err)
	}

	g.Notes = g.note + note_seperator + string(statusJson)

	return g.Task, nil
}

func (t *googleTask) GetId() string {
	return t.Id
}

func (t *googleTask) GetTitle() string {
	return t.Title
}

func (t *googleTask) SetTitle(title string) Task {
	t.Title = title
	return t
}

func (t *googleTask) GetNote() string {
	return t.note
}

func (t *googleTask) SetNote(note string) Task {
	t.note = note
	return t
}

func (t *googleTask) GetFocus() bool {
	return t.Focus
}

func (t *googleTask) SetFocus(focus bool) Task {
	t.Focus = focus
	return t
}

func (t *googleTask) GetStatus() Status {
	switch t.Status {
	case gtaskStatusDone:
		return StatusDone
	case gtaskStatusTodo:
		return StatusTodo
	default:
		return StatusInvalid
	}
}

func (t *googleTask) SetStatus(status Status) Task {
	switch status {
	case StatusDone:
		t.Status = gtaskStatusDone
	case StatusTodo:
		t.Status = gtaskStatusTodo
	default:
		t.Status = gtaskStatusInvalid
	}

	return t
}

func (t *googleTask) GetProject() string {
	return t.Project
}

func (t *googleTask) SetProject(project string) Task {
	t.Project = project
	return t
}

func (t *googleTask) GetTags() []string {
	return t.Tags
}

func (t *googleTask) SetTag(tag string) Task {
	t.Tags = append(t.Tags, tag)
	return t
}

func (t *googleTask) GetCategory() Category {
	return Category(t.Category)
}

func (t *googleTask) SetCategory(category Category) Task {
	t.Category = category
	return t
}

func (t *googleTask) GetDue() string {
	return t.Due
}

func (t *googleTask) SetDue(due string) Task {
	t.Due = due
	return t
}

func (t *googleTask) GetCompleted() string {
	if t.Completed == nil {
		return ""
	}

	return *t.Completed
}

func (t *googleTask) SetCompleted(c string) Task {
	t.Completed = &c
	return t
}

func (t *googleTask) GetUpdated() string {
	return t.Updated
}

func (t *googleTask) Error() error {
	return t.err
}
