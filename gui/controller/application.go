package controller

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/dustinliu/taskcommander/core"
	"github.com/dustinliu/taskcommander/gui/view"
	"github.com/dustinliu/taskcommander/service"
)

var categories = []string{"Inbox", "Next", "Someday", "Focus"}

type Application struct {
	fyne.App
	service service.TaskService
	config  core.Config
}

func NewApplication() (*Application, error) {
	s, err := service.NewService()
	if err != nil {
		return nil, fmt.Errorf("failed to create service: %w", err)
	}
	if err := s.Init(); err != nil {
		return nil, fmt.Errorf("failed to init service: %w", err)
	}

	app := &Application{
		App:     app.New(),
		service: s,
		config:  core.GetConfig(),
	}

	app.Settings().SetTheme(view.NewMyTheme())

	return app, nil
}

func (app *Application) Run() {
	win := app.NewWindow(core.AppName)
	win.SetMaster()
	win.Resize(fyne.NewSize(800, 600))

	left := container.NewGridWithRows(1, view.NewCategoryList(categories))
	right := container.NewVSplit(view.NewTaskList(), view.NewTaskList())
	root := container.NewHSplit(left, right)
	root.SetOffset(0.2)
	win.SetContent(root)

	win.ShowAndRun()
}
