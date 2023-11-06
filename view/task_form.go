package view

import (
	"github.com/dustinliu/taskcommander/service"
	"github.com/rivo/tview"
)

const (
	DescLabel    = "Desciption"
	ProjectLabel = "Project"
	Notelabel    = "Note"
)

func newTaskForm(task *service.Task, onSave func(*service.Task), onCancel func()) *tview.Form {
	form := tview.NewForm()

	s := func() {
		task := &service.Task{
			Description: form.GetFormItemByLabel(DescLabel).(*tview.InputField).GetText(),
			Project:     form.GetFormItemByLabel(ProjectLabel).(*tview.InputField).GetText(),
		}
		onSave(task)
	}

	form.AddInputField(DescLabel, "", 0, nil, nil).
		AddInputField(ProjectLabel, "", 0, nil, nil).
		AddTextArea(Notelabel, "", 0, 0, 0, nil).
		AddButton("Save", s).AddButton("Cancel", onCancel).
		SetBorder(true)

	return form
}
