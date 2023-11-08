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

type TaskForm struct {
	*tview.Form

	projects []string
	tags     []string
}

func newTaskForm(task *service.Task, onSave func(*service.Task), onCancel func()) *TaskForm {
	form := &TaskForm{
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
		AddFormItem(newInputField(TagsLabel, comp(form.tags), nil)).
		AddFormItem(newInputField(ProjectLabel, comp(form.projects), nil)).
		AddTextArea(Notelabel, "", 0, 0, 0, nil).
		AddButton("Save", save).AddButton("Cancel", onCancel).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyEsc:
				onCancel()
				return nil
			}
			return event
		}).
		SetBorder(true)

	return form
}

func comp(candiates []string) func(string) []string {
	return func(input string) []string {
		list := []string{}
		for _, candidate := range candiates {
			if input != "" && strings.HasPrefix(strings.ToLower(candidate), strings.ToLower(input)) {
				list = append(list, candidate)
			}
		}
		return list
	}
}

func newInputField(label string,
	comp func(string) []string,
	comped func(text string, index int, source int) bool) *tview.InputField {
	mainStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)
	selectStyle := tcell.StyleDefault.Background(tcell.ColorLightGreen).Foreground(tcell.ColorBlack)
	input := tview.NewInputField().
		SetLabel(label).
		SetFieldWidth(0).
		SetAutocompleteStyles(tcell.ColorWhite, mainStyle, selectStyle).
		SetAutocompleteFunc(comp).SetAutocompletedFunc(comped)

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
