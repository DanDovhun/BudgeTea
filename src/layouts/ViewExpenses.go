package layouts

import (
	"BudgeTea/datamng"
	"BudgeTea/maths"
	"BudgeTea/report"
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

// Finds a month
func pastMonthReport(root fyne.App, home fyne.Window) {
	// Create a new window
	window := root.NewWindow("Month and Year - BudgeTea")

	// Create a title label
	title := widget.NewLabel("Enter a month and year")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	// Entries for mont and a year
	month := widget.NewEntry()
	year := widget.NewEntry()

	// Set placeholders
	month.SetPlaceHolder("Month")
	year.SetPlaceHolder("Year")

	// Fill the window with content
	window.SetContent(container.NewVBox(
		title,
		month,
		year,

		// Submits the input
		widget.NewButton("Submit", func() {
			Popup(root, window, "To be implemented", false)
		}),
	))

	// Sets the window's size and shows it
	window.Resize(fyne.NewSize(400, 150))
	window.Show()
}

// Report creation options
func createReport(root fyne.App, home fyne.Window) {
	// Create a new window window
	window := root.NewWindow("Report Options - BudgeTea")

	// Fill the content
	window.SetContent(container.NewVBox(
		// Export current month's spending report
		widget.NewButton("This month's spending", func() {
			err := report.ExportCurrentExpenses()

			if err != nil {

			}
		}),

		// Export spending report from any previous month
		widget.NewButton("Past months' spending", func() {
			pastMonthReport(root, window) // Gathers the month and exports the report
		}),

		// Exports the entire spending history
		widget.NewButton("Whole spending history", func() {
			Popup(root, window, "To be implemented", false)
		}),
	))

	// Set size and show the window
	window.Resize(fyne.NewSize(400, 120))
	window.Show()
}

// A winodow to view summary of expenses and to export a report
func ViewExpensesLayout(root fyne.App, home fyne.Window) {
	// Create a new window, set it to master and hide the home window
	window := root.NewWindow("Add New Expense - BudgeTea")
	window.SetMaster()
	home.Hide()

	// Get current month and year
	month := time.Now().Month().String()
	year := time.Now().Year()

	// Creates a title label, allignes it to center and sets it to bold
	title := widget.NewLabel(fmt.Sprintf("Expenses for: %v %v", month, year))
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	info, err := report.GetThisMonthsExpenses()

	if err != nil {
		Popup(root, home, "Cannot get expenses", true)
		return
	}

	currency, err := datamng.GetCurrency()

	if err != nil {
		Popup(root, home, "Cannot get currency", true)
		return
	}

	var spending string

	if info.Budget > info.TotalSpending {
		under := info.Budget - info.TotalSpending
		percentage := (under / info.Budget) * 100
		spending = fmt.Sprintf("%v %v less than the budget (%v", maths.Round(info.Budget-info.TotalSpending, 2),
			currency, maths.Round(percentage, 2)) + "% under budget)"
	} else {
		over := info.TotalSpending - info.Budget
		percentage := (over / info.Budget) * 100
		spending = fmt.Sprintf("%v %v above the budget (%v", maths.Round(over, 2),
			currency, maths.Round(percentage, 2)) + "% above budget)"
	}

	// Fill the window with the content
	window.SetContent(container.NewVBox(
		// Add title
		title,

		// Displays the budget, money spet that month spending and whether the user went over the budget
		widget.NewLabel(fmt.Sprintf("Budget: %v %v", info.Budget, currency)),
		widget.NewLabel(fmt.Sprintf("Money spent: %v %v", maths.Round(info.TotalSpending, 2), currency)),
		widget.NewLabel(spending),

		// Display spending by category
		widget.NewLabel(fmt.Sprintf("Grocieries: %v %v", maths.Round(info.Groceries, 2), currency)),
		widget.NewLabel(fmt.Sprintf("Hobbies: %v %v", maths.Round(info.Hobbies, 2), currency)),
		widget.NewLabel(fmt.Sprintf("Rent: %v %v", maths.Round(info.Rent, 2), currency)),
		widget.NewLabel(fmt.Sprintf("Travel: %v %v", maths.Round(info.Travel, 2), currency)),
		widget.NewLabel(fmt.Sprintf("Miscelanious: %v %v", maths.Round(info.Miscelanious, 2), currency)),
		widget.NewLabel(fmt.Sprintf("Other bills: %v %v", maths.Round(info.OtherBills, 2), currency)),

		// Option to create a spending report
		widget.NewButton("Create a report", func() {
			createReport(root, window)
		}),

		// Go home
		widget.NewButton("Home", func() {
			home.SetMaster() // Set home to master
			home.Show()      // Show the home window
			window.Hide()    // Hide the previous window
		}),
	))

	window.Resize(fyne.NewSize(600, 200))
	window.Show()
}
