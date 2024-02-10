package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/dustinliu/taskcommander/service"
)

type TaskList struct {
	*widget.Table
	tasks []service.Task
}

func NewTaskList() *TaskList {
	t := &TaskList{}
	table := widget.NewTable(t.length, createCell, t.updateCell)

	t.Table = table
	t.tasks = []service.Task{}
	return t
}

func (t *TaskList) length() (row, col int) {
	return len(t.tasks), 1
}

func createCell() fyne.CanvasObject {
	label := widget.NewLabel("")
	return label
}

func (t *TaskList) updateCell(id widget.TableCellID, cell fyne.CanvasObject) {
	cell.(*widget.Label).SetText(t.tasks[id.Row].GetTitle())
}

func (t *TaskList) SetTasks(tasks []service.Task) {
	t.tasks = tasks
	t.Refresh()
}
