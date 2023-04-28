// BudgeTea, (C) Daniel Dovhun
// Simple budgetting tool built with Go

package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"

	"BudgeTea/layouts"
)

func main() { // Main function
	root := app.New()                  // Create an application instance
	home := root.NewWindow("BudgeTea") // Create a home window

	label := widget.NewLabel("BudgeTea")         // Create a title
	label.Alignment = fyne.TextAlignCenter       // Allign to center
	label.TextStyle = fyne.TextStyle{Bold: true} // Make the title bold

	homeLayout := container.NewVBox( // Create a container
		label, // Add title label

		widget.NewButton("Add an Expense", func() { // Create a button that switches to ExpenseAdditionWindow
			go layouts.ExpenseAdditionWindow(root, home) // Replace home window with the new one
		}),
		widget.NewButton("Expense Report", func() { // Create a button that switches to ViewExpensesLayout
			go layouts.ViewExpensesLayout(root, home) // Switch to the new layout
		}),
		widget.NewButton("Preferences", func() { // Create a button that switches to Preferences
			go layouts.Preferences(root, home) // Switch to preferences
		}),
	)

	home.SetContent(homeLayout) // Set content to the current container

	home.ShowAndRun() // Start the application and show the window
}
