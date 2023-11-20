package view

import (
	"github.com/dustinliu/taskcommander/service"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	mainLabel  = "main"
	modalLabel = "modal"
)

type TUI struct {
	*tview.Application
	pages          *tview.Pages
	categoryPannel *categoryPannel
	taskPannel     *taskPannel
	infoPannel     *infoPannel
	messageLine    *messageLine
}

func NewTUI() *TUI {
	tui := &TUI{
		tview.NewApplication(),
		tview.NewPages(),

		newCategoryPannel(),
		newTaskPannel(),
		newInfoPannel(),
		newMessageLine(),
	}

	tui.pages.AddPage(mainLabel,
		newMainWin(tui.categoryPannel, tui.taskPannel, tui.infoPannel, tui.messageLine),
		true, true)
	tui.SetRoot(tui.pages, true)
	tui.SetInputCapture(tui.inputHandler)
	return tui
}

func newMainWin(cat *categoryPannel,
	task *taskPannel,
	info *infoPannel,
	msg *messageLine) *tview.Flex {
	flex := tview.NewFlex().
		AddItem(cat, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(task, 0, 1, false).
			AddItem(info, 0, 1, false), 0, 4, false)

	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(flex, 0, 1, true).
		AddItem(msg, 1, 1, false)
}

func newModal(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewGrid().
		SetColumns(0, width, 0).
		SetRows(0, height, 0).
		AddItem(p, 1, 1, 1, 1, 0, 0, true)
}

func (tui *TUI) inputHandler(event *tcell.EventKey) *tcell.EventKey {
	if tui.categoryPannel.HasFocus() || tui.taskPannel.HasFocus() {
		switch event.Key() {
		case tcell.KeyTab:
			tui.ChangeFocus()
		case tcell.KeyEsc:
			service.Events <- service.NewEventQuit()
		case tcell.KeyRune:
			switch event.Rune() {
			case 'q':
				service.Events <- service.NewEventQuit()
			case 'a':
				tui.OpenTaskForm(nil)
			}
		}
	}
	return event
}

func (tui *TUI) OpenTaskForm(task *service.Task) {
	onSave := func(task *service.Task) {
		service.Events <- service.NewEventAddTask(*task)
		tui.pages.RemovePage(modalLabel)
	}

	onCancel := func() {
		tui.pages.RemovePage(modalLabel)
	}

	tui.pages.AddPage(modalLabel,
		newModal(newTaskForm(task, onSave, onCancel), 60, 35),
		true, true)
}

func (tui *TUI) PrintMessage(msg string) {
	tui.messageLine.setText(msg)
}

func (tui *TUI) SetTaskList(tasks []service.Task) {
	tui.taskPannel.setTasks(tasks)
}

func (tui *TUI) SetTaskInfo(task *service.Task) {
	tui.infoPannel.SetTask(task)
}

func (tui *TUI) GetCurrentCategory() service.Category {
	row, _ := tui.categoryPannel.GetSelection()
	return service.Category(row)
}

func (tui *TUI) ChangeFocus() {
	if tui.categoryPannel.HasFocus() {
		tui.SetFocus(tui.taskPannel)
	} else {
		tui.SetFocus(tui.categoryPannel)
	}
}
