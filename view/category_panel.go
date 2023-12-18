package view

import (
	"github.com/dustinliu/taskcommander/service"
	"github.com/rivo/tview"
)

type categoryPannel struct {
	*tview.Table
}

func newCategoryPannel() *categoryPannel {
	var pannel *categoryPannel
	table := tview.NewTable().SetSelectable(true, false)
	table.SetBorder(true).SetBorderPadding(0, 0, 1, 1).
		SetFocusFunc(func() { table.SetBorderStyle(focusStyle) }).
		SetBlurFunc(func() { table.SetBorderStyle(blurStyle) })

	table.SetCellSimple(int(service.Inbox), 0, "Inbox").
		SetCellSimple(int(service.Next), 0, "Next").
		SetCellSimple(int(service.Someday), 0, "Someday").
		SetCellSimple(int(service.Focus), 0, "Focus").
		Select(int(service.Next), 0).
		SetSelectionChangedFunc(func(row, _ int) {
			service.Events <- service.NewEventCategoryChange(service.Category(row))
		})

	pannel = &categoryPannel{
		table,
	}

	return pannel
}
