// BudgeTea, (C) Daniel Dovhun
// Simple budgetting tool built with Go
// Version 0.0.2

package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"

	"BudgeTea/layouts"
)

func main() {
	root := app.New()
	home := root.NewWindow("BudgeTea")

	label := widget.NewLabel("BudgeTea")
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle = fyne.TextStyle{Bold: true}

	homeLayout := container.NewVBox(
		label,

		widget.NewButton("Add an Expense", func() {
			go layouts.ExpenseAdditionWindow(root, home)
		}),
		widget.NewButton("Expense Report", func() {
			layouts.ViewExpensesLayout(root, home)
		}),
		widget.NewButton("Preferences", func() {
			layouts.Preferences(root, home)
		}),
	)

	home.SetContent(homeLayout)

	home.Resize(fyne.NewSize(400, 200))
	home.ShowAndRun()
}
