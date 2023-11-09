package view

import "github.com/rivo/tview"

type messageLine struct {
	*tview.TextView
}

func newMessageLine() *messageLine {
	return &messageLine{
		tview.NewTextView(),
	}
}

func (l *messageLine) setText(text string) {
	l.TextView.SetText(text)
}
