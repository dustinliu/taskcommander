package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/dustinliu/taskcommander/core"
)

func createDialog(title, msg, buttonText string, onClose func(), win fyne.Window) *dialog.CustomDialog {
	label := widget.NewLabel(msg)
	label.Wrapping = fyne.TextWrapWord
	scroll := container.NewVScroll(label)
	errDialog := dialog.NewCustom(title, buttonText, scroll, win)
	factor := float32(6)
	winSize := win.Content().Size()
	core.GetLogger().Debugf("winSize: %+v", winSize)
	errDialog.SetOnClosed(onClose)
	errDialog.Resize(fyne.NewSize(winSize.Width*factor, winSize.Height*factor))
	errDialog.Show()
	return errDialog
}

func ShowError(err error, fatal bool, onClose func(), win fyne.Window) {
	var title, buttonText string
	if fatal {
		title = "Fatal Error"
		buttonText = "Quit"
	} else {
		title = "Error"
		buttonText = "OK"
	}
	d := createDialog(title, err.Error(), buttonText, onClose, win)
	d.Show()
}
