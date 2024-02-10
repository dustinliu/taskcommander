package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	xtheme "fyne.io/x/fyne/theme"
)

type MyTheme struct {
	fyne.Theme
}

func DefaultTheme() *MyTheme {
	return &MyTheme{
		Theme: xtheme.AdwaitaTheme(),
	}
}

func (t MyTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	// case theme.SizeNameInnerPadding:
	// return 3
	// case theme.SizeNamePadding:
	// return 5
	case theme.SizeNameSeparatorThickness:
		return 0
	default:
		return t.Theme.Size(name)
	}
}
