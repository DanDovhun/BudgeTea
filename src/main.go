// BudgeTea, (C) Daniel Dovhun
// Simple budgetting tool built with Go

package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"

	"BudgeTea/forex"
	"BudgeTea/layouts"
)

func main() { // Main function
	_, notConnected := forex.Convert("SEK", "USD", 100)

	root := app.New()                  // Create an application instance
	home := root.NewWindow("BudgeTea") // Create a home window

	if notConnected != nil {
		layouts.Popup(root, home, "Internet connection not found.\n\nPlease make sure to be connected to the internet \nto get full BudgeTea experience", true)
	}

	label := widget.NewLabel("BudgeTea")         // Create a title
	label.Alignment = fyne.TextAlignCenter       // Allign to center
	label.TextStyle = fyne.TextStyle{Bold: true} // Make the title bold

	homeLayout := container.NewVBox( // Create a container
		label, // Add title label

		widget.NewButton("Add an Expense", func() { // Create a button that switches to ExpenseAdditionWindow
			layouts.ExpenseAdditionWindow(root, home) // Replace home window with the new one
		}),
		widget.NewButton("Expense Report", func() { // Create a button that switches to ViewExpensesLayout
			layouts.ViewExpensesLayout(root, home) // Switch to the new layout
		}),
		widget.NewButton("Preferences", func() { // Create a button that switches to Preferences
			layouts.Preferences(root, home) // Switch to preferences
		}),
	)

	home.SetContent(homeLayout) // Set content to the current container
	home.Resize(fyne.NewSize(250, 50))
	home.ShowAndRun() // Start the application and show the window
}
