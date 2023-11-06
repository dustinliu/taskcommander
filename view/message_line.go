package view

import "github.com/rivo/tview"

type MessageLine struct {
	*tview.TextView
}

func newMessageLine() *MessageLine {
	return &MessageLine{
		tview.NewTextView(),
	}
}

func (l *MessageLine) setText(text string) {
	l.TextView.SetText(text)
}
