package view

import (
	"strconv"

	"github.com/dustinliu/taskcommander/service"
	"github.com/rivo/tview"
)

type InfoPannel struct {
	*tview.Table
}

func newInfoPannel() *InfoPannel {
	var pannel *InfoPannel
	table := tview.NewTable().SetSelectable(false, false)
	table.SetBorder(true).SetBorderPadding(0, 0, 1, 1).
		SetFocusFunc(func() { table.SetBorderStyle(FocusStyle) }).
		SetBlurFunc(func() { table.SetBorderStyle(BlurStyle) })

	pannel = &InfoPannel{
		table,
	}

	return pannel
}

func (p *InfoPannel) SetTask(task *service.Task) {
	if task == nil {
		return
	}
	p.Clear()

	row := 0
	p.SetCellSimple(row, 0, "Id").SetCellSimple(0, 1, strconv.Itoa(task.Id))

	row++
	p.SetCellSimple(row, 0, "Description").SetCellSimple(row, 1, task.Description)

	if len(task.Tags) > 0 {
		row++
		tags := ""
		for _, tag := range task.Tags {
			tags += tag + " "
		}
		p.SetCellSimple(row, 0, "Tags").SetCellSimple(row, 1, tags)
	}

	if task.Project != "" {
		row++
		p.SetCellSimple(row, 0, "Project").SetCellSimple(row, 1, task.Project)
	}

	row++
	p.SetCellSimple(row, 0, "UpdateAt").SetCellSimple(row, 1, task.UpdatedAt.Format("2006-01-02 15:04:05"))

	row++
	p.SetCellSimple(row, 0, "CreateAt").SetCellSimple(row, 1, task.CreatedAt.Format("2006-01-02 15:04:05"))
}
