package view

import (
	"strings"

	"github.com/dustinliu/taskcommander/service"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	DescLabel    = "Desciption"
	ProjectLabel = "Project"
	TagsLabel    = "Tags"
	Notelabel    = "Note"
)

type taskForm struct {
	*tview.Form

	projects []string
	tags     []string
}

func newTaskForm(
	task service.Task,
	projects []string,
	tags []string,
	onSave func(service.Task),
	onCancel func(),
) *taskForm {
	form := &taskForm{
		tview.NewForm(),
		projects,
		tags,
	}

	save := func() {
		task := task.SetTitle(form.GetFormItemByLabel(DescLabel).(*tview.InputField).GetText()).
			SetProject(form.GetFormItemByLabel(ProjectLabel).(*projectInputField).GetText())
		onSave(task)
	}

	form.AddInputField(DescLabel, "", 0, nil, nil).
		AddFormItem(newProjectInputField(projects)).
		AddFormItem(newTagsInputField(tags)).
		AddTextArea(Notelabel, "", 0, 0, 0, nil).
		AddButton("Save", save).AddButton("Cancel", onCancel).
		SetCancelFunc(onCancel).
		SetBorder(true)

	return form
}

type projectInputField struct {
	*tview.InputField
	projects []string
}

func newProjectInputField(projects []string) *projectInputField {
	p := &projectInputField{}
	p.InputField = newInputField(ProjectLabel, p.comp, nil, nil)
	p.projects = projects

	return p
}

func (p *projectInputField) comp(input string) []string {
	list := []string{}
	for _, project := range p.projects {
		if input != "" && strings.HasPrefix(strings.ToLower(project), strings.ToLower(input)) {
			list = append(list, project)
		}
	}
	return list
}

type tagsInputField struct {
	*tview.InputField
	tags     []string
	compList []string
	changed  bool
}

func newTagsInputField(tags []string) *tagsInputField {
	t := &tagsInputField{}
	t.InputField = newInputField(TagsLabel, t.complete, t.completed, t.onChange)
	t.tags = tags
	t.compList = []string{}
	t.changed = true

	return t
}

func (t *tagsInputField) complete(input string) []string {
	keyword := t.GetKeywords(input)
	if keyword == "" {
		return []string{}
	}

	if t.changed {
		t.compList = []string{}
		for _, tags := range t.tags {
			if strings.HasPrefix(strings.ToLower(tags), strings.ToLower(keyword)) {
				t.compList = append(t.compList, tags)
			}
		}
	}
	return t.compList
}

func (t *tagsInputField) completed(text string, index int, source int) bool {
	if text == "" {
		return true
	}

	content := t.GetText()
	tags := strings.Fields(content)
	if len(tags) > 0 {
		tags = tags[:len(tags)-1]
	}
	tags = append(tags, text)
	t.SetText(strings.Join(tags, " "))

	switch source {
	case tview.AutocompletedNavigate:
		t.changed = false
		return false
	case tview.AutocompletedEnter:
		t.changed = false
		return true
	case tview.AutocompletedTab:
		t.changed = false
		return true
	default:
		return true
	}
}

func (t *tagsInputField) onChange(text string) {
	t.changed = true
}

func (t *tagsInputField) GetKeywords(content string) string {
	if len(content) == 0 || content[len(content)-1] == ' ' {
		return ""
	}
	tags := strings.Fields(content)
	if len(tags) > 0 {
		return tags[len(tags)-1]
	}
	return ""
}

func newInputField(label string,
	complete func(string) []string,
	completed func(string, int, int) bool,
	onchanged func(string),
) *tview.InputField {
	input := tview.NewInputField().
		SetLabel(label).
		SetFieldWidth(0).
		SetAutocompleteStyles(tcell.ColorWhite, listMainStyle, listSelectStyle).
		SetAutocompleteFunc(complete)
	if completed != nil {
		input.SetAutocompletedFunc(completed)
	}

	if onchanged != nil {
		input.SetChangedFunc(onchanged)
	}

	input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlN:
			return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
		case tcell.KeyCtrlP:
			return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
		}
		return event
	})
	return input
}
