package view

import "github.com/rivo/tview"

type CmdLine struct {
	*tview.InputField
}

func newCmdLine() *CmdLine {
	return &CmdLine{
		tview.NewInputField(),
	}
}

func (cl *CmdLine) setText(text string) {
	cl.InputField.SetText(text)
}

func (cl *CmdLine) clear() {
	cl.SetText("")
}
