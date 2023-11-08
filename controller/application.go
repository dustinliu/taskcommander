package controller

import (
	"sync"

	"github.com/dustinliu/taskcommander/service"
	"github.com/dustinliu/taskcommander/view"
	"github.com/fsnotify/fsnotify"
	"github.com/gdamore/tcell/v2"
)

var app *Application
var once sync.Once

type Application struct {
	tui *view.TUI
}

func GetApplication() *Application {
	once.Do(func() {
		app = &Application{
			tui: view.NewTUI(),
		}
	})

	return app
}

func (app *Application) Run() error {
	go app.handlerEvent()
	service.Events <- service.NewEventCategoryChange(service.Next)

	service.GetLogger().Info("======================= application started ==================")
	return app.tui.Run()
}

func (app *Application) Stop() {
	app.tui.Stop()
	service.GetLogger().Info("======================= application stopped ==================")
}

func (app *Application) handlerEvent() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		app.tui.PrintMessage(err.Error())
	}
	defer watcher.Close()
	watcher.Add(GetConfig().Data_location)

	for {
		select {
		case event := <-service.Events:
			app.handlerViewEvent(event)
		case event := <-watcher.Events:
			app.handleWatcherEvent(event)
		case err := <-watcher.Errors:
			service.GetLogger().Debug("watcher error: ", err)
			app.tui.PrintMessage(err.Error())
		}
	}
}

func (app *Application) handlerViewEvent(event tcell.Event) {
	service.GetLogger().Debug("view event received: ", event)
	switch event := event.(type) {
	case service.EventCategoryChange:
		app.tui.QueueUpdateDraw(func() { app.categoryChanged(event.Category) })
	case service.EventTaskChange:
		app.tui.QueueUpdateDraw(func() { app.taskChanged(event.Task) })
	case service.EventQuit:
		app.Stop()
	case service.EventAddTask:
		app.tui.QueueUpdateDraw(func() { app.addTask(event.Task) })
	}
}

func (app *Application) handleWatcherEvent(event fsnotify.Event) {
	service.GetLogger().Debug("watcher event received: ", event)
	switch event.Op {
	case fsnotify.Write:
		app.tui.QueueUpdateDraw(func() { app.refresh() })
	}
}

func (app *Application) addTask(task service.Task) {
	task.Category = app.tui.GetCurrentCategory()
	if err := service.AddTask(&task); err != nil {
		app.tui.PrintMessage(err.Error())
		service.GetLogger().Error(err.Error())
		return
	}
}

func (app *Application) categoryChanged(cat service.Category) {
	tasks, err := service.ListTasksByCategory(cat)
	if err != nil {
		app.tui.PrintMessage(err.Error())
		service.GetLogger().Error(err.Error())
		return
	}
	app.tui.SetTaskList(tasks)
}

func (app *Application) taskChanged(task service.Task) {
	app.tui.SetTaskInfo(&task)
}

func (app *Application) refresh() {
	app.categoryChanged(app.tui.GetCurrentCategory())
}
