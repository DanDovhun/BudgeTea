package layouts

import (
	"BudgeTea/datamng"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

// User enters information about a new expense
func ExpenseAdditionWindow(root fyne.App, home fyne.Window) {
	// Stores user's choice for denomination and category
	var denomination, category string

	// Create a new window, set it to master and hide the home window
	window := root.NewWindow("Add New Expense - BudgeTea")
	window.SetMaster()
	home.Hide()

	// Creates a title label, allignes it to center and sets it to bold
	label := widget.NewLabel("BudgeTea")
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle = fyne.TextStyle{Bold: true}

	// Create an entry for the user to enter purchase title
	expenseTitle := widget.NewEntry()
	expenseTitle.SetPlaceHolder("Purchase title:")

	// Create an entry for the user to enter the purchase's cost
	expenseCost := widget.NewEntry()
	expenseCost.SetPlaceHolder("Purchase cost:")

	// Denomination options
	denoms := widget.NewRadioGroup([]string{
		"EUR",
		"SEK",
		"USD",
	}, func(choice string) {
		denomination = choice
	})

	// Category options
	categories := widget.NewRadioGroup([]string{
		"Groceries",
		"Hobbies",
		"Rent",
		"Other bills",
		"Travel",
		"Miscelanious",
	}, func(choice string) {
		category = choice
	})

	// Set content
	window.SetContent(container.NewVBox(
		// Add title and entry fields
		label,
		expenseTitle,
		expenseCost,

		// Add denomination otpions
		widget.NewLabel("Denomination"),
		denoms,

		// Add category options
		widget.NewLabel("Categories"),
		categories,

		// Submit input
		widget.NewButton("Submit", func() {
			// If the user didn't enter expense's title
			if len(expenseTitle.Text) == 0 {
				Popup(root, window, "Please enter the expense's name", true)

				return
			}

			// If the user didn't enter expense's cost
			if len(expenseCost.Text) == 0 {
				Popup(root, window, "Please enter expense's cost", true)

				return
			}

			// If the user didn't choose a category
			if len(category) == 0 {
				Popup(root, window, "Pleasxe select a category", true)

				return
			}

			// If the user didn't choose a denomination
			if len(denomination) == 0 {
				Popup(root, window, "Please select a denomination", false)

				return
			}

			// Try to convert cost input into a float64
			cost, err := strconv.ParseFloat(expenseCost.Text, 64)

			// If it cannot be converted
			if err != nil {
				// Send an error message
				Popup(root, window, "Please enter the expense's cost as a number", true)

				return
			}

			// Create a new expense object
			expense := datamng.NewExpense(expenseTitle.Text, category, denomination, cost)

			// Try to add the expense into the database
			err = expense.Add(expense)

			// If it cannot be added
			if err != nil {
				// Send an error message
				Popup(root, window, "Couldn't add expense", true)

				return
			}

			// If everything goes right, send a success message
			Popup(root, window, "Expense added succesfully", false)
		}),

		widget.NewButton("Home", func() {
			home.SetMaster() // Set home to master
			home.Show()      // Show the home window
			window.Hide()    // Hide the previous window
		}),
	))

	window.Resize(fyne.NewSize(600, 200))
	window.Show()
}
