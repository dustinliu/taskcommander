package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/dustinliu/taskcommander/event"
	"github.com/dustinliu/taskcommander/service"
)

type CategoryPannel struct {
	*widget.List
	categories []service.Category
}

func NewCategoryPannel(categories []service.Category, onSelect func(widget.ListItemID)) *CategoryPannel {
	p := &CategoryPannel{}
	list := widget.NewList(p.length, createLabel, p.updateLabel)
	p.List = list
	p.categories = categories
	if onSelect != nil {
		p.OnSelected = onSelect
	} else {
		p.OnSelected = func(id widget.ListItemID) {
			event.QueueEvent(event.NewEventCategoryChanged(p.categories[id]))
		}
	}

	return p
}

func (p *CategoryPannel) length() int {
	return len(p.categories)
}

func createLabel() fyne.CanvasObject {
	label := canvas.NewText("", color.White)
	label.TextStyle = fyne.TextStyle{Bold: true}
	return label
}

func (c *CategoryPannel) updateLabel(i widget.ListItemID, label fyne.CanvasObject) {
	l := label.(*canvas.Text)
	l.Text = c.categories[i].Name()
}
