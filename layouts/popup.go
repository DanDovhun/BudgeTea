package layouts

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func error(root fyne.App, home fyne.Window, message string, isError bool) {
	window := root.NewWindow("")

	if isError {
		window.SetTitle("Error has occured - BudgeTea")
	} else {
		window.SetTitle("Message - BudgeTea")
	}

	window.SetContent(container.NewVBox(
		widget.NewLabel(message),
	))

	window.Resize(fyne.NewSize(150, 100))
	window.Show()
}
