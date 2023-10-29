package view

import "github.com/rivo/tview"

type TaskPannel struct {
	*tview.Table
}

func newTaskPannel() *TaskPannel {
	var pannel *TaskPannel
	table := tview.NewTable().SetSelectable(true, true)
	table.SetBorder(true)
	table.SetCellSimple(0, 0, "aaaaaa").
		SetCellSimple(1, 0, "bbbbbb").
		SetCellSimple(2, 0, "cccccc").
		SetCellSimple(3, 0, "dddddd").
		SetCellSimple(4, 0, "eeeee")

	pannel = &TaskPannel{
		table,
	}

	return pannel
}
