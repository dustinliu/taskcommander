package controller

import (
	"sync"

	"github.com/dustinliu/taskcommander/service"
	"github.com/dustinliu/taskcommander/view"
	"github.com/gdamore/tcell/v2"
)

var app *Application
var once sync.Once

type Application struct {
	tui *view.TaskTUI
}

func GetApplication() *Application {
	once.Do(func() {
		app = &Application{
			tui: view.NewTaskTUI(),
		}
		app.tui.SetInputCapture(app.inputHandler)
	})

	return app
}

// TODO move to tui?
func (app *Application) inputHandler(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyTab:
		app.tui.ChangeFocus()
	case tcell.KeyRune:
		switch event.Rune() {
		case 'q':
			app.tui.Stop()
		}
	}
	return event
}

func (app *Application) Run() error {
	go app.eventHandler()
	view.SendEvent(view.NewEventCategoryChange(app.tui.GetCurrentCategory()))

	return app.tui.Run()
}

func (app *Application) eventHandler() {
	for {
		switch event := view.PollEvent().(type) {
		case view.EventCategoryChange:
			app.tui.QueueUpdateDraw(func() {
				app.changeCategory(event)
			})
		}
	}
}

func (app *Application) changeCategory(event view.EventCategoryChange) {
	tasks, err := service.ListTasks(event.Category)
	if err != nil {
		app.tui.PrintMessage(err.Error())
		return
	}
	app.tui.SetTasks(tasks)
}
