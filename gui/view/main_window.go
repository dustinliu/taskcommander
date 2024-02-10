package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/dustinliu/taskcommander/core"
)

type MainWindow struct {
	fyne.Window

	app fyne.App
}

func NewMainWindow(app fyne.App, left, right fyne.CanvasObject) *MainWindow {
	main := &MainWindow{
		Window: app.NewWindow(core.AppName),
		app:    app,
	}

	main.SetMaster()
	main.Resize(fyne.NewSize(800, 600))

	root := container.NewHSplit(container.NewGridWithRows(1, left), right)
	root.SetOffset(0.2)
	main.SetContent(root)

	return main
}
