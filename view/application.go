package view

import (
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *Application
var once sync.Once

type Application struct {
	tvapp      *tview.Application
	mainPannel *MainPannel
	taskPannel *TaskPannel
}

func init() {
	app = &Application{
		tview.NewApplication(),
		newMainPannel(),
		newTaskPannel(),
	}

	flex := tview.NewFlex().
		AddItem(app.mainPannel, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(app.taskPannel, 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 0, 1, false), 0, 4, false)

	root := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(flex, 0, 1, true).
		AddItem(tview.NewTextArea(), 1, 1, false)

	app.tvapp.SetRoot(root, true).SetFocus(app.mainPannel).SetInputCapture(eventHandler)
}

func (app *Application) Run() error {
	return app.tvapp.Run()
}

func GetApplication() *Application {
	return app
}

func eventHandler(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyTab:
		changeFocus()
	}
	return event
}

func changeFocus() {
	if app.mainPannel.HasFocus() {
		app.tvapp.SetFocus(app.taskPannel)
	} else {
		app.tvapp.SetFocus(app.mainPannel)
	}
}
