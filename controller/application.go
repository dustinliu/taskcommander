package controller

import (
	"github.com/dustinliu/taskcommander/gtask"
	"github.com/dustinliu/taskcommander/logger"
	"github.com/dustinliu/taskcommander/service"
	"github.com/dustinliu/taskcommander/view"
	"github.com/fsnotify/fsnotify"
	"github.com/gdamore/tcell/v2"
)

type Application struct {
	tui     *view.TUI
	service service.TaskService
}

func NewApplication() *Application {
	s, err := gtask.NewGoogleTaskService()
	if err != nil {
		panic(err)
	}
	app := &Application{
		tui:     view.NewTUI(),
		service: s,
	}

	return app
}

// TODO: refactor event
func (app *Application) Run() error {
	go app.handlerEvent()
	// service.Events <- service.NewEventCategoryChange(service.Next)

	logger.GetLogger().Info("======================= application started ==================")
	return app.tui.Run()
}

func (app *Application) Stop() {
	app.tui.Stop()
	logger.GetLogger().Info("======================= application stopped ==================")
}

func (app *Application) handlerEvent() {
	// TODO: refactor event
	for {
		select {
		// case event := <-service.Events:
		// app.handlerViewEvent(event)
		}
	}
}

// TODO: refactor event
func (app *Application) handlerViewEvent(event tcell.Event) {
	logger.GetLogger().Debug("view event received: ", event)
	//switch event := event.(type) {
	// case service.EventCategoryChange:
	// app.tui.QueueUpdateDraw(func() { app.categoryChanged(event.Category) })
	// case service.EventTaskChange:
	// app.tui.QueueUpdateDraw(func() { app.taskChanged(event.Task) })
	// case service.EventQuit:
	// app.Stop()
	// case service.EventAddTask:
	// app.tui.QueueUpdateDraw(func() { app.addTask(event.Task) })
	//}
}

func (app *Application) handleWatcherEvent(event fsnotify.Event) {
	logger.GetLogger().Debug("watcher event received: ", event)
	switch event.Op {
	case fsnotify.Write:
		app.tui.QueueUpdateDraw(func() { app.refresh() })
	}
}

// TODO: refactor category
func (app *Application) addTask(task service.Task) {
	// task.Category = app.tui.GetCurrentCategory()
	if err := app.service.AddTask(task); err != nil {
		app.tui.PrintMessage(err.Error())
		logger.GetLogger().Error(err.Error())
		return
	}
}

// TODO: refactor cagegory
//func (app *Application) categoryChanged(cat service.Category) {
//tasks, err := service.ListTasksByCategory(cat)
//if err != nil {
//app.tui.PrintMessage(err.Error())
//service.GetLogger().Error(err.Error())
//return
//}
//app.tui.SetTaskList(tasks)
//}

func (app *Application) taskChanged(task service.Task) {
	app.tui.SetTaskInfo(task)
}

// TODO*: refactor cagegory
func (app *Application) refresh() {
	// app.categoryChanged(app.tui.GetCurrentCategory())
}
