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
			// To be implemented
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
			// To be implemented
		}),

		// Export spending report from any previous month
		widget.NewButton("Past months' spending", func() {
			pastMonthReport(root, window) // Gathers the month and exports the report
		}),

		// Exports the entire spending history
		widget.NewButton("Whole spending history", func() {
			// To be implemented
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

	// Fill the window with the content
	window.SetContent(container.NewVBox(
		// Add title
		title,

		// Displays the budget, money spet that month spending and whether the user went over the budget
		widget.NewLabel("Budget: "),
		widget.NewLabel("Money spent: "),
		widget.NewLabel("n less then the budget (p%)"),

		// Display spending by category
		widget.NewLabel("Grocieries: "),
		widget.NewLabel("Hobbies: "),
		widget.NewLabel("Rent: "),
		widget.NewLabel("Other bills: "),
		widget.NewLabel("Travel: "),
		widget.NewLabel("Miscelanious: "),

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
