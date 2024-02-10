package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MyTheme struct {
	fyne.Theme
}

func NewMyTheme() *MyTheme {
	return &MyTheme{
		Theme: theme.DefaultTheme(),
	}
}

func (t MyTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameInnerPadding:
		return 0
	case theme.SizeNamePadding:
		return 5
	case theme.SizeNameSeparatorThickness:
		return 0
	default:
		return t.Theme.Size(name)
	}
}
