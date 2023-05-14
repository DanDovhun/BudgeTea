package layouts

import (
	"BudgeTea/datamng"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Gathers user input for new budget
func setBudget(root fyne.App, home fyne.Window) {
	window := root.NewWindow("Set budget - BudgeTea") // Create the window and set its title

	budget, _ := datamng.GetBudget()
	currency, _ := datamng.GetCurrency()
	budgetTitle := widget.NewLabel(fmt.Sprintf("Current budget: %v %v", budget, currency))

	budgetEntry := widget.NewEntry()         // Create a new entry field
	budgetEntry.SetPlaceHolder("New Budget") // Set the placeholder to 'New Budget'

	window.SetContent(container.NewVBox( // Creates a new container and sets the content
		budgetTitle,
		budgetEntry, // Adds the entry field

		widget.NewButton("Set budget", func() { // Button to set the new budget
			budget, err := strconv.ParseFloat(budgetEntry.Text, 64)

			if err != nil {
				Popup(root, home, "Please only enter numbers", true)
				return
			}

			err = datamng.SetBudget(budget)

			if err != nil {
				Popup(root, window, "Error has occured", true)

				return
			}

			Popup(root, home, "Budget set", false)
			window.Hide()
		}),
	))

	window.Resize(fyne.NewSize(300, 100)) // Resizes to 100x100
	window.Show()                         // Shows the window
}

// Gathers user input to set a new preferred denomination
func setDenomination(root fyne.App, home fyne.Window) string {
	var denomination string // Stores user's choice

	window := root.NewWindow("Set denomination - BudgeTea") // New window

	currency, errs := datamng.GetCurrency() // Gets currenctly set denomination
	titleMessage := ""

	if errs != nil { // If the program cannot get currently set denomination
		Popup(root, window, "Cannot get current denomination", true)
		return ""
	}

	// Set the title
	titleMessage = fmt.Sprintf("Set denomination (current = %v)", currency)

	title := widget.NewLabel(titleMessage)       // Title label, will also show currently set denomination
	title.Alignment = fyne.TextAlignCenter       // Allign it to center
	title.TextStyle = fyne.TextStyle{Bold: true} // Make it bold

	denoms := widget.NewRadioGroup([]string{ // Create an instance of a group of radio buttons
		"EUR", // Euros
		"SEK", // Swedish Kronar
		"USD", // American Dollars
	}, func(choice string) { // Gets the user input
		denomination = choice // Store it in the denomination variable
	})

	window.SetContent(container.NewVBox( // Create a new container and make it window's content
		title,
		denoms,

		widget.NewButton("Set denomination", func() {
			// If no denomination is selected
			if len(denomination) == 0 {
				Popup(root, window, "Please select a denomination", true)

				return
			}

			datamng.SetCurrency(denomination)
			Popup(root, window, "Denomination set succesfully", false)
			window.Close()
		}),
	))

	window.Resize(fyne.NewSize(320, 100))
	window.Show()

	return denomination
}

// Preferences window
func Preferences(root fyne.App, home fyne.Window) {
	// Create a new window, set it to master and hide the home window
	window := root.NewWindow("Preferences - BudgeTea")
	window.SetMaster()
	home.Hide()

	// Creates a title label, allignes it to center and makes it bold
	title := widget.NewLabel("Preferences")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	// Fill the window with a new container
	window.SetContent(container.NewVBox(
		title, // Add title

		// Opens tne setBudget window
		widget.NewButton("Set Budget", func() {
			setBudget(root, window)
		}),

		// Opens tne setDenomination window
		widget.NewButton("Change preferred denomination", func() {
			setDenomination(root, window)
		}),

		// Goes back to window
		widget.NewButton("Home", func() {
			home.SetMaster() // Set home to master
			home.Show()      // Show the home window
			window.Hide()    // Hide the previous window
		}),
	))

	// Resizes to 300x100 and shows the window
	window.Resize(fyne.NewSize(300, 100))
	window.Show()
}
