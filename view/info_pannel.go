package view

import (
	"strconv"

	"github.com/dustinliu/taskcommander/service"
	"github.com/rivo/tview"
)

type InfoPannel struct {
	*tview.Table
}

func newInfoPannel() *InfoPannel {
	var pannel *InfoPannel
	table := tview.NewTable().SetSelectable(false, false)
	table.SetBorder(true).SetBorderPadding(0, 0, 1, 1).
		SetFocusFunc(func() { table.SetBorderStyle(FocusStyle) }).
		SetBlurFunc(func() { table.SetBorderStyle(BlurStyle) })

	pannel = &InfoPannel{
		table,
	}

	return pannel
}

func (p *InfoPannel) SetTask(task *service.Task) {
	if task == nil {
		return
	}
	p.Clear()

	i := 0
	p.SetCellSimple(i, 0, "Id").SetCellSimple(0, 1, strconv.Itoa(task.Id))

	i++
	p.SetCellSimple(i, 0, "Description").SetCellSimple(1, 1, task.Description)

	if task.Project != "" {
		i++
		p.SetCellSimple(i, 0, "Project").SetCellSimple(2, 1, task.Project)
	}
}
