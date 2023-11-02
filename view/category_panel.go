package view

import (
	"github.com/dustinliu/taskcommander/service"
	"github.com/rivo/tview"
)

type CategoryPannel struct {
	*tview.Table
}

func newCategoryPannel() *CategoryPannel {
	var pannel *CategoryPannel
	table := tview.NewTable().SetSelectable(true, false)
	table.SetBorder(true).SetBorderPadding(0, 0, 1, 0).
		SetFocusFunc(func() { table.SetBorderStyle(FocusStyle) }).
		SetBlurFunc(func() { table.SetBorderStyle(BlurStyle) })

	table.SetCellSimple(int(service.Inbox), 0, "Inbox").
		SetCellSimple(int(service.Next), 0, "Next").
		SetCellSimple(int(service.Someday), 0, "Someday").
		SetCellSimple(int(service.Focus), 0, "Focus").
		Select(int(service.Next), 0).
		SetSelectionChangedFunc(func(row, colum int) {
			SendEvent(NewEventCategoryChange(service.Category(row)))
		})

	pannel = &CategoryPannel{
		table,
	}

	return pannel
}
