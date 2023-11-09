package view

import (
	"strings"
	"unicode"

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

func newTaskForm(task *service.Task, onSave func(*service.Task), onCancel func()) *taskForm {
	form := &taskForm{
		tview.NewForm(),
		service.ListProjects(),
		service.ListTags(),
	}

	save := func() {
		task := &service.Task{
			Description: form.GetFormItemByLabel(DescLabel).(*tview.InputField).GetText(),
			Project:     form.GetFormItemByLabel(ProjectLabel).(*tview.InputField).GetText(),
		}
		onSave(task)
	}

	form.AddInputField(DescLabel, "", 0, nil, nil).
		AddFormItem(newProjectInputField()).
		AddFormItem(newTagsInputField()).
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

func newProjectInputField() *projectInputField {
	p := &projectInputField{}
	p.InputField = newInputField(ProjectLabel, p.comp, nil)
	p.projects = service.ListProjects()

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

func newTagsInputField() *tagsInputField {
	t := &tagsInputField{}
	t.InputField = newInputField(TagsLabel, t.complete, t.completed)
	t.tags = service.ListTags()
	t.compList = []string{}
	t.changed = true

	return t
}

func (t *tagsInputField) complete(input string) []string {
	if t.changed {
		t.compList = []string{}
		for _, tags := range t.tags {
			keyword := t.GetKeywords()
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

	switch source {
	case tview.AutocompletedNavigate:
		t.changed = false
		return false
	case tview.AutocompletedEnter:
		t.SetText(strings.Join(tags, " "))
		t.changed = true
		return true
	case tview.AutocompletedTab:
		return true
	default:
		return true
	}
}

func (t *tagsInputField) GetKeywords() string {
	tags := strings.Fields(t.GetText())
	r := []rune(tags[len(tags)-1])
	if len(tags) == 0 || unicode.IsSpace(r[len(r)-1]) {
		return ""
	}
	if len(tags) == 1 {
		return tags[0]
	}

	return tags[len(tags)-1]
}

func newInputField(label string,
	complete func(string) []string,
	completed func(text string, index int, source int) bool) *tview.InputField {
	input := tview.NewInputField().
		SetLabel(label).
		SetFieldWidth(0).
		SetAutocompleteStyles(tcell.ColorWhite, listMainStyle, listSelectStyle).
		SetAutocompleteFunc(complete)
	if completed != nil {
		input.SetAutocompletedFunc(completed)
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
