package view

import (
	"github.com/dustinliu/taskcommander/service"
	"github.com/rivo/tview"
)

type infoPannel struct {
	*tview.Table
}

func newInfoPannel() *infoPannel {
	var pannel *infoPannel
	table := tview.NewTable().SetSelectable(false, false)
	table.SetBorder(true).SetBorderPadding(0, 0, 1, 1).
		SetFocusFunc(func() { table.SetBorderStyle(focusStyle) }).
		SetBlurFunc(func() { table.SetBorderStyle(blurStyle) })

	pannel = &infoPannel{
		table,
	}

	return pannel
}

func (p *infoPannel) SetTask(task service.Task) {
	if task == nil {
		return
	}
	p.Clear()

	row := 0
	p.SetCellSimple(row, 0, "Id").SetCellSimple(0, 1, task.GetId())

	row++
	p.SetCellSimple(row, 0, "Description").SetCellSimple(row, 1, task.GetTitle())

	if len(task.GetTags()) > 0 {
		row++
		tags := ""
		for _, tag := range task.GetTags() {
			tags += tag + " "
		}
		p.SetCellSimple(row, 0, "Tags").SetCellSimple(row, 1, tags)
	}

	if task.GetProject() != "" {
		row++
		p.SetCellSimple(row, 0, "Project").SetCellSimple(row, 1, task.GetProject())
	}
}
