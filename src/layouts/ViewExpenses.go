package layouts

import (
	"BudgeTea/datamng"
	"BudgeTea/maths"
	"BudgeTea/report"
	"fmt"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
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
	monthInp := widget.NewEntry()
	yearInp := widget.NewEntry()

	// Set placeholders
	monthInp.SetPlaceHolder("Month")
	yearInp.SetPlaceHolder("Year")

	// Fill the window with content
	window.SetContent(container.NewVBox(
		title,
		monthInp,
		yearInp,

		// Submits the input
		widget.NewButton("Submit", func() {
			var month time.Month

			yearInt, err := strconv.ParseInt(yearInp.Text, 64, 10)

			if err != nil {
				Popup(root, window, err.Error(), true)
				return
			}

			year := int(yearInt)

			switch monthInp.Text {
			case "January", "january", "Jan", "jan", "1", "01":
				month = 1

			case "February", "february", "Feb", "feb", "2", "02":
				month = 2

			case "March", "march", "Mar", "mar", "3", "03":
				month = 3

			case "April", "april", "Apr", "apr", "4", "04":
				month = 4

			case "May", "may", "5", "05":
				month = 5

			case "June", "june", "Jun", "jun", "6", "06":
				month = 6

			case "July", "july", "Jul", "jul", "7", "07":
				month = 7

			case "August", "august", "Aug", "aug", "8", "08":
				month = 8

			case "September", "september", "Sep", "sep", "9", "09":
				month = 9

			case "October", "october", "Oct", "oct", "10":
				month = 10

			case "November", "november", "Nov", "nov", "11":
				month = 11

			case "December", "december", "Dec", "dec", "12":
				month = 12

			default:
				Popup(root, window, "Invalid month", true)
				return
			}

			folderBrowser := root.NewWindow("Export expense report")
			fileTitle := fmt.Sprintf("%v-%v_Expense_Report", month, year)

			fileName := widget.NewEntry()
			fileName.SetText(fileTitle)
			fileName.SetPlaceHolder("File name")

			folderBrowser.SetContent(container.NewVBox(
				fileName,

				widget.NewButton("Save as CSV", func() {
					window.Resize(fyne.NewSize(600, 400))
					folderBrowser.Hide()
					name := fileName.Text

					if len(name) == 0 {
						Popup(root, folderBrowser, "Please enter file's name", true)

						return
					}

					if !strings.Contains(name, ".csv") {
						name += ".csv"
					}

					var save *dialog.FileDialog = new(dialog.FileDialog)

					save = dialog.NewFileSave(func(w fyne.URIWriteCloser, e error) {
						window.Hide()
						if w == nil {
							return
						}

						if e != nil {
							return
						}

						data, err := datamng.GetData()

						if err != nil {
							Popup(root, window, "Unexpected error", true)

							return
						}

						index, month := data.GetMonth(month, year)

						if index == -1 {
							Popup(root, window, "Month not found", true)

							return
						}

						normal, err := report.NormaliseExpenses(month, data.Denomination)

						if err != nil {
							Popup(root, window, "Unexpected error", true)

							return
						}

						list := report.CreateOneMonthReport(normal, data.Denomination)

						w.Write([]byte(list))

						Popup(root, home, fmt.Sprintf("Succesfully exported the report to %v", w.URI()), false)

						window.Resize(fyne.NewSize(400, 120))
					}, window)

					save.SetFileName(name)
					save.Show()
				}),
			))

			folderBrowser.Resize(fyne.NewSize(350, 100))
			folderBrowser.Show()
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
			folderBrowser := root.NewWindow("Export expense report")

			year := time.Now().Local().Year()
			month := time.Now().Local().Month()

			fileTitle := fmt.Sprintf("%v-%v_Expense_Report", month, year)

			fileName := widget.NewEntry()
			fileName.SetText(fileTitle)
			fileName.SetPlaceHolder("File name")

			folderBrowser.SetContent(container.NewVBox(
				fileName,

				widget.NewButton("Save as csv", func() {
					window.Resize(fyne.NewSize(600, 400))
					folderBrowser.Hide()
					name := fileName.Text

					if len(name) == 0 {
						Popup(root, folderBrowser, "Please enter file's name", true)

						return
					}

					if !strings.Contains(name, ".csv") {
						name += ".csv"
					}

					var save *dialog.FileDialog = new(dialog.FileDialog)

					save = dialog.NewFileSave(func(w fyne.URIWriteCloser, e error) {
						window.Hide()
						if w == nil {
							return
						}

						if e != nil {
							return
						}

						data, err := datamng.GetData()

						if err != nil {
							Popup(root, window, "Unexpected error", true)

							return
						}

						month := data.GetLastMonth()

						normal, err := report.NormaliseExpenses(month, data.Denomination)

						if err != nil {
							Popup(root, window, "Unexpected error", true)

							return
						}

						list := report.CreateOneMonthReport(normal, data.Denomination)

						w.Write([]byte(list))

						Popup(root, home, fmt.Sprintf("Succesfully exported the report to %v", w.URI()), false)

						window.Resize(fyne.NewSize(400, 120))
					}, window)

					save.SetFileName(name)
					save.Show()
				}),

				/*
					widget.NewButton("Save as excel", func() {
						window.Resize(fyne.NewSize(600, 400))
						folderBrowser.Hide()

						name := fileName.Text

						if strings.Contains(name, ".csv") {
							name = strings.Replace(name, ".csv", ".xlsx", 1)
						} else if !strings.Contains(name, ".xlsx") {
							name += ".xlsx"
						}

						var save *dialog.FileDialog = new(dialog.FileDialog)

						save = dialog.NewFileSave(func(w fyne.URIWriteCloser, e error) {
							window.Hide()
							if w == nil {
								return
							}

							if e != nil {
								return
							}

							data, err := datamng.GetData()

							if err != nil {
								Popup(root, window, "Unexpected error", true)

								return
							}

							month := data.GetLastMonth()

							normal, err := report.NormaliseExpenses(month, data.Denomination)

							if err != nil {
								Popup(root, window, "Unexpected error", true)

								return
							}

							report.CreateLastMonthExcel(w, normal, data.Budget, data.Denomination, name)
						}, window)

						save.SetFileName(name)
						save.Show()
					}),
				*/
			))

			folderBrowser.Resize(fyne.NewSize(350, 100))
			folderBrowser.Show()
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
