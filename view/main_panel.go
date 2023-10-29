package view

import (
	"github.com/rivo/tview"
)

const (
	Inbox = iota
	Space
	Next
	Someday
	Focus
)

type MainPannel struct {
	*tview.Table
}

func newMainPannel() *MainPannel {
	var pannel *MainPannel
	table := tview.NewTable().SetSelectable(true, true)
	table.SetBorder(true)
	table.SetCellSimple(Inbox, 0, "Inbox").
		SetCellSimple(Next, 0, "Next").
		SetCellSimple(Someday, 0, "Someday").
		SetCellSimple(Focus, 0, "Focus")

	table.Select(2, 0).SetSelectionChangedFunc(pannel.onChange)
	pannel = &MainPannel{
		table,
	}

	return pannel
}

func (pannel *MainPannel) onChange(row, column int) {
	switch row {
	case Inbox:

	}
}
