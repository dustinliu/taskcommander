package view

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/dustinliu/taskcommander/service"
	"github.com/stretchr/testify/assert"
)

func TestCategoryPannel(t *testing.T) {
	defer test.NewApp()
	var cat service.Category
	list := NewCategoryPannel(service.Categories(), func(id widget.ListItemID) {
		cat = service.Categories()[id]
	})
	window := test.NewWindow(list)
	defer window.Close()
	window.Resize(list.MinSize().Max(fyne.NewSize(150, 200)))

	list.Select(0)
	assert.Equal(t, cat, service.CategoryInbox)

	list.Select(1)
	assert.Equal(t, cat, service.CategoryNext)

	list.Select(2)
	assert.Equal(t, cat, service.CategorySomeday)

	list.Select(3)
	assert.Equal(t, cat, service.CategoryFocus)
}
