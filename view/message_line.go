package view

import "github.com/rivo/tview"

type MessageLine struct {
	*tview.InputField
}

func newMessageLine() *MessageLine {
	return &MessageLine{
		tview.NewInputField(),
	}
}

func (l *MessageLine) setText(text string) {
	l.clear()
	l.InputField.SetText(text)
}

func (l *MessageLine) clear() {
	l.SetText("")
}
