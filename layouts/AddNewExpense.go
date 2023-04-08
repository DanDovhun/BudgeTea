package layouts

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

// User enters information about a new expense
func ExpenseAdditionWindow(root fyne.App, home fyne.Window) {
	var denomination, category string

	window := root.NewWindow("Add New Expense - Budgetie")
	window.SetMaster()
	home.Hide()

	label := widget.NewLabel("Budgetie")
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle = fyne.TextStyle{Bold: true}

	expenseTitle := widget.NewEntry()
	expenseTitle.SetPlaceHolder("Expense name:")

	expenseCost := widget.NewEntry()
	expenseCost.SetPlaceHolder("Expense cost:")

	denoms := widget.NewRadioGroup([]string{
		"EUR",
		"SEK",
		"USD",
	}, func(choice string) {
		denomination = choice
	})

	categories := widget.NewRadioGroup([]string{
		"Grocieries",
		"Hobbies",
		"Rent",
		"Other bills",
		"Travel",
		"Miscelanious",
	}, func(choice string) {
		category = choice
	})

	window.SetContent(container.NewVBox(
		label,
		expenseTitle,
		expenseCost,

		widget.NewLabel("Denomination"),

		denoms,

		widget.NewLabel("Categories"),

		categories,

		widget.NewButton("Submit", func() {
			fmt.Println(denomination)
			fmt.Println(category)
		}),

		widget.NewButton("Home", func() {
			home.SetMaster()
			home.Show()
			window.Hide()
		}),
	))
	window.Resize(fyne.NewSize(400, 200))
	window.Show()
}
