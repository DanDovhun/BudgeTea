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
	home := root.NewWindow("Budgetie")

	label := widget.NewLabel("Budgetie")
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle = fyne.TextStyle{Bold: true}

	home.SetContent(container.NewVBox(
		label,

		widget.NewButton("Add an Expense", func() {
			layouts.ExpenseAdditionWindow(root, home)
		}),
		widget.NewButton("Expense Report", func() {}),
		widget.NewButton("Preferences", func() {}),
	))

	home.Resize(fyne.NewSize(400, 200))
	home.ShowAndRun()
}
