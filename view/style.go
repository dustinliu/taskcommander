package view

import "github.com/gdamore/tcell/v2"

var (
	focusStyle      = tcell.StyleDefault.Foreground(tcell.ColorDarkCyan)
	blurStyle       = tcell.StyleDefault
	listMainStyle   = tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)
	listSelectStyle = tcell.StyleDefault.Background(tcell.ColorLightGreen).Foreground(tcell.ColorBlack)
)
