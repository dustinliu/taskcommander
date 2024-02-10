package view

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/dustinliu/taskcommander/core"
	"github.com/dustinliu/taskcommander/service"
)

type TaskPannel struct {
	*widget.Table
	tasks []service.Task
}

func NewTaskPannel() *TaskPannel {
	t := &TaskPannel{}

	table := widget.NewTable(t.length, createCell, t.updateCell)
	t.Table = table

	return t
}

func (t *TaskPannel) length() (row, col int) {
	r := len(t.tasks)
	return r, 1
}

func createCell() fyne.CanvasObject {
	label := widget.NewLabel("")
	return label
}

func (t *TaskPannel) updateCell(id widget.TableCellID, cell fyne.CanvasObject) {
	switch id.Col {
	case 0:
		cell.(*widget.Label).SetText(t.tasks[id.Row].GetTitle())
	case 1:
		cell.(*widget.Label).SetText(strings.Join(t.tasks[id.Row].GetTags(), ", "))
	}
}

func (t *TaskPannel) SetTasks(tasks []service.Task) {
	core.GetLogger().Debugf("set tasks: %v, length: %d\n", tasks, len(tasks))
	t.tasks = tasks
	t.Refresh()
}
