package layouts

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

// Finds a month
func pastMonthReport(root fyne.App, home fyne.Window) {
	window := root.NewWindow("Month and Year - BudgeTea")

	title := widget.NewLabel("Enter a month and year")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	month := widget.NewEntry()
	year := widget.NewEntry()

	month.SetPlaceHolder("Month")
	year.SetPlaceHolder("Year")

	window.SetContent(container.NewVBox(
		title,
		month,
		year,

		widget.NewButton("Submit", func() {}),
	))

	window.Show()
}

// Report creation options
func createReport(root fyne.App, home fyne.Window) {
	window := root.NewWindow("Report Options - BudgeTea")

	window.SetContent(container.NewVBox(
		widget.NewButton("This month's spending", func() {}),
		widget.NewButton("Past months' spending", func() {
			pastMonthReport(root, window)
		}),
		widget.NewButton("Whole spending history", func() {}),
	))

	window.Show()
}

// A winodow to view summary of expenses and to export a report
func ViewExpensesLayout(root fyne.App, home fyne.Window) {
	window := root.NewWindow("Add New Expense - BudgeTea")
	window.SetMaster()
	home.Hide()

	month := time.Now().Month().String()
	year := time.Now().Year()

	title := widget.NewLabel(fmt.Sprintf("Expenses for: %v %v", month, year))
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	window.SetContent(container.NewVBox(
		title,

		widget.NewLabel("Budget: "),
		widget.NewLabel("Money spent: "),
		widget.NewLabel("n less then the budget (p%)"),
		widget.NewLabel("Grocieries: "),
		widget.NewLabel("Hobbies: "),
		widget.NewLabel("Rent: "),
		widget.NewLabel("Other bills: "),
		widget.NewLabel("Travel: "),
		widget.NewLabel("Miscelanious: "),

		widget.NewButton("Create a report", func() {
			createReport(root, window)
		}),

		widget.NewButton("Back", func() {
			home.SetMaster()
			home.Show()
			window.Hide()
		}),
	))

	window.Show()
}
