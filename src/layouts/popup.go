package layouts

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// A popup window in case something happens (error, warning...)
// Besides the app instance and a window it takes a message and whether the
// popup happens due to an error or not
func Popup(root fyne.App, home fyne.Window, message string, isError bool) {
	// Create the window
	window := root.NewWindow("")

	if isError { // If popup appears due to an error
		window.SetTitle("Error has occured - BudgeTea")
	} else { // Otherwise
		window.SetTitle("Message - BudgeTea")
	}

	// Create an instance of the label with the message and allign it to center
	msg := widget.NewLabel(message)
	msg.Alignment = fyne.TextAlignCenter

	// Show a message
	window.SetContent(container.NewVBox(
		msg,
	))

	// Set size and show the window
	window.Resize(fyne.NewSize(300, 100))
	window.Show()
}
