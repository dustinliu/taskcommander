package view

import (
	"github.com/dustinliu/taskcommander/service"
	"github.com/rivo/tview"
)

type TaskPannel struct {
	*tview.Table
}

func newTaskPannel() *TaskPannel {
	table := tview.NewTable().SetSelectable(true, false)
	table.SetBorder(true).SetBorderPadding(0, 0, 1, 0)

	table.SetFocusFunc(func() { table.SetBorderStyle(FocusStyle) }).
		SetBlurFunc(func() { table.SetBorderStyle(BlurStyle) })

	return &TaskPannel{
		table,
	}
}

func (p *TaskPannel) setTasks(tasks []service.Task) {
	p.Clear()
	for i, task := range tasks {
		p.SetCell(i, 0, tview.NewTableCell(task.Description).SetReference(task))
	}
}
