package view

import (
	"github.com/dustinliu/taskcommander/service"
	"github.com/rivo/tview"
)

type TaskTUI struct {
	*tview.Application
	categoryPannel *CategoryPannel
	taskPannel     *TaskPannel
	messageLine    *MessageLine
}

func NewTaskTUI() *TaskTUI {
	tui := &TaskTUI{
		tview.NewApplication(),
		newCategoryPannel(),
		newTaskPannel(),
		newMessageLine(),
	}

	flex := tview.NewFlex().
		AddItem(tui.categoryPannel, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tui.taskPannel, 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 0, 1, false), 0, 4, false)

	root := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(flex, 0, 1, true).
		AddItem(tview.NewTextArea(), 1, 1, false)
	tui.Application.SetRoot(root, true)
	return tui
}

func (tui *TaskTUI) PrintMessage(msg string) {
	tui.messageLine.setText(msg)
}

func (tui *TaskTUI) SetTasks(tasks []service.Task) {
	tui.taskPannel.setTasks(tasks)
}

func (tui *TaskTUI) GetCurrentCategory() service.Category {
	row, _ := tui.categoryPannel.GetSelection()
	return service.Category(row)
}

// TODO check if this is the right way to do this
func (tui *TaskTUI) ChangeFocus() {
	if tui.categoryPannel.HasFocus() {
		tui.Application.SetFocus(tui.taskPannel)
	} else {
		tui.Application.SetFocus(tui.categoryPannel)
	}
}
