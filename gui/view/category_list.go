package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type CategoryList struct {
	*widget.List
	categories []string
}

func NewCategoryList(categories []string) fyne.CanvasObject {
	p := &CategoryList{}
	list := widget.NewList(p.length, createLabel, p.updateLabel)
	p.List = list
	p.categories = categories
	return p
}

func (p *CategoryList) length() int {
	return len(p.categories)
}

func createLabel() fyne.CanvasObject {
	label := widget.NewLabel("")
	label.TextStyle.Bold = true
	return label
}

func (c *CategoryList) updateLabel(i widget.ListItemID, label fyne.CanvasObject) {
	label.(*widget.Label).SetText(c.categories[i])
}
