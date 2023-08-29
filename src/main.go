/*
* Name: BudgeTea
* Author: Daniel Dovhun
* Description: A simple budget managing tool built with GoLang and Fyne.io
 */

package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"BudgeTea/datamng"
	"BudgeTea/forex"
	"BudgeTea/layouts"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func fillup() {
	fl, err := os.Open("datamng_test/test_cases.txt")

	if err != nil {
		return
	}

	bytes, err := ioutil.ReadAll(fl)

	arr := strings.Split(string(bytes), "\n")

	for _, i := range arr {
		ex := strings.Split(i, ",")

		price, _ := strconv.ParseFloat(ex[1], 64)
		data, _ := datamng.NewExpense(ex[0], ex[2], price)
		data.Add(data)
	}
}

func main() { // Main function
	fillup()
	_, notConnected := forex.Convert("SEK", "USD", 100)

	root := app.New()                  // Create an application instance
	home := root.NewWindow("BudgeTea") // Create a home window

	home.SetFixedSize(true) // Make the window

	if notConnected != nil {
		layouts.Popup(root, home, "Internet connection not found.\n\nPlease make sure to be connected to the internet \nto get full BudgeTea experience", true)
	}

	label := widget.NewLabel("BudgeTea")         // Create a title
	label.Alignment = fyne.TextAlignCenter       // Allign to center
	label.TextStyle = fyne.TextStyle{Bold: true} // Make the title bold

	homeLayout := container.NewVBox( // Create a container
		label, // Add title label

		widget.NewButton("Add an Expense", func() { // Create a button that switches to ExpenseAdditionWindow
			layouts.ExpenseAdditionWindow(root, home) // Replace home window with the new one
		}),

		widget.NewButton("Expense Report", func() { // Create a button that switches to ViewExpensesLayout
			layouts.ViewExpensesLayout(root, home) // Switch to the new layout
		}),
		widget.NewButton("Preferences", func() { // Create a button that switches to Preferences
			layouts.Preferences(root, home) // Switch to preferences
		}),
	)

	home.SetContent(homeLayout) // Set content to the current container
	home.Resize(fyne.NewSize(250, 50))
	home.ShowAndRun() // Start the application and show the window
}
