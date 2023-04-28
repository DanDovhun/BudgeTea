package layouts

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

// Gathers user input for new budget
func setBudget(root fyne.App, home fyne.Window) {
	window := root.NewWindow("Set budget - BudgeTea") // Create the window and set its title

	budgetEntry := widget.NewEntry()         // Create a new entry field
	budgetEntry.SetPlaceHolder("New Budget") // Set the placeholder to 'New Budget'

	window.SetContent(container.NewVBox( // Creates a new container and sets the content
		budgetEntry, // Adds the entry field

		widget.NewButton("Set budget", func() { // Button to set the new budget
			// To be implemented
		}),
	))

	window.Resize(fyne.NewSize(100, 100)) // Resizes to 100x100
	window.Show()                         // Shows the window
}

// Gathers user input to set a new preferred denomination
func setDenomination(root fyne.App, home fyne.Window) {
	var denomination string // Stores user's choice

	window := root.NewWindow("Set denomination - BudgeTea") // New window

	title := widget.NewLabel("Denominations (current = SEK): ") // Title label, will also show currently set denomination
	title.Alignment = fyne.TextAlignCenter                      // Allign it to center
	title.TextStyle = fyne.TextStyle{Bold: true}                // Make it bold

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
			fmt.Println(denomination)

			// To be implemented
		}),
	))
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

	Popup(root, home, "Hello world", false)

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