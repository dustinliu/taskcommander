package view

import (
	"github.com/dustinliu/taskcommander/service"
	"github.com/rivo/tview"
)

type taskPannel struct {
	*tview.Table
}

func newTaskPannel() *taskPannel {
	table := tview.NewTable().SetSelectable(true, false)
	table.SetBorder(true).SetBorderPadding(0, 0, 1, 1)

	table.SetFocusFunc(func() { table.SetBorderStyle(focusStyle) }).
		SetBlurFunc(func() { table.SetBorderStyle(blurStyle) })

	table.SetSelectionChangedFunc(func(row, colum int) {
		t := table.GetCell(row, colum).GetReference()
		if t != nil {
			// TODO: refactor category
			// service.Events <- service.NewEventTaskChange(t.(service.Task))
		}
	})

	return &taskPannel{
		table,
	}
}

func (p *taskPannel) setTasks(tasks []service.Task) {
	p.Clear()
	for i, task := range tasks {
		p.SetCell(i, 0, tview.NewTableCell(task.GetTitle()).SetReference(task))
	}
	p.Select(0, 0)
}

func (p *taskPannel) GetSelectedTask() (*service.Task, bool) {
	row, _ := p.GetSelection()
	task := p.GetCell(row, 0).GetReference()
	if task == nil {
		return nil, false
	}

	t := task.(service.Task)
	return &t, true
}
