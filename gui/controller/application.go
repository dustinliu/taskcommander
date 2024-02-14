package controller

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/dustinliu/taskcommander/core"
	"github.com/dustinliu/taskcommander/event"
	"github.com/dustinliu/taskcommander/gui/view"
	"github.com/dustinliu/taskcommander/service"
)

type Application struct {
	fyne.App
	service service.TaskService

	mainWindow     fyne.Window
	categoryPannel *view.CategoryPannel
	taskPannel     *view.TaskPannel
}

func NewApplication() *Application {
	cat := view.NewCategoryPannel(service.Categories(), nil)
	task := view.NewTaskPannel()
	ap := app.New()
	app := &Application{
		App:            ap,
		categoryPannel: cat,
		taskPannel:     task,
		mainWindow:     view.NewMainWindow(ap, cat, task),
	}

	lifecycle := app.Lifecycle()
	lifecycle.SetOnStarted(app.init)

	lifecycle.SetOnStopped(func() {
		event.Wg.Wait()
	})

	app.Settings().SetTheme(view.DefaultTheme())

	return app
}

func (app *Application) Run() {
	core.GetLogger().Infof("debug: %v\n", core.GetConfig().Debug)
	event.ListenEvent(app.inputHandler)
	app.mainWindow.ShowAndRun()
}

func (app *Application) init() {
	s, err := service.NewService()
	if err != nil {
		event.SendEvent(event.NewEventError(err, true))
		return
	}
	app.service = s

	// TODO: implement url handler
	err = s.OAuth2(func(u string) error {
		url, err := url.Parse(u)
		if err != nil {
			return err
		}
		err = app.OpenURL(url)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		event.SendEvent(event.NewEventError(err, true))
		return
	}

	app.categoryPannel.Select(0)
}

func (app *Application) inputHandler(e event.Event) event.Event {
	switch e := e.(type) {
	case event.EventCategoryChanged:
		app.categoryChangeHandler(e.Category)
		core.GetLogger().Debugf("category(%v) changed, get tasks\n", e.Category)
	case event.EventError:
		var onClose func()
		if e.Fatal {
			onClose = app.Quit
		}
		view.ShowError(fmt.Errorf("error: %w", e.Err), e.Fatal, onClose, app.mainWindow)
	}

	return e
}

// TODO: implement error handling
func (app *Application) categoryChangeHandler(c service.Category) {
	tasks, err := app.service.ListTodoTasks()
	core.GetLogger().Debugf("tasks length: %d\n", len(tasks))
	if err != nil {
		view.ShowError(fmt.Errorf("failed to handle category change: %w", err), false, nil, app.mainWindow)
		return
	}
	app.taskPannel.SetTasks(tasks)
}
