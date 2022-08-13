package main

import (
	"BudgetManager/db/data"
	"BudgetManager/db/expenses"
	"BudgetManager/funcs"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ErrorMessage(a fyne.App, h fyne.Window, msg string) {
	win := a.NewWindow("Error has occured")

	win.SetContent(container.NewVBox(
		widget.NewLabel(msg),
		widget.NewButton("Ok", func() {
			win.Hide()
		}),
	))
	win.Show()
}

func AddExpenses(a fyne.App, h fyne.Window) {
	win := a.NewWindow("Add an expense")
	h.Hide()
	win.SetMaster()

	var category string

	title := widget.NewLabel("Add an expense")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	price := widget.NewEntry()
	price.SetPlaceHolder("Price (SEK)")

	cats := widget.NewLabel("Categories:")
	cats.Alignment = fyne.TextAlign(fyne.OrientationHorizontalRight)

	categories := widget.NewRadioGroup([]string{"Groceries", "Travel", "Others"}, func(value string) {
		category = value
	})

	win.SetContent(container.NewVBox(
		title,
		price,
		cats,
		categories,

		widget.NewButton("Submit", func() {
			prc, err := strconv.ParseFloat(price.Text, 64)

			if err != nil && len(price.Text) > 0 {
				ErrorMessage(a, win, "Please only enter numbers")
			}

			if len(price.Text) == 0 {
				ErrorMessage(a, win, "Please enter price")
			}

			if prc < 0 && len(price.Text) != 0 {
				ErrorMessage(a, win, "Price cannot be negative")
			}

			if len(category) == 0 {
				ErrorMessage(a, win, "Please select a category")
			}

			err = data.Add(prc, category)

			if err != nil {
				ErrorMessage(a, win, err.Error())
			}
		}),

		widget.NewButton("Home", func() {
			win.Hide()
			h.Show()
			h.SetMaster()
		}),
	))

	win.Resize(fyne.NewSize(300, 300))
	win.Show()
}

func SeeExpenses(a fyne.App, h fyne.Window) {
	title := "Expenses"
	win := a.NewWindow(title)

	h.Hide()
	win.SetMaster()

	exp, err := expenses.GetThisMonthsExpenses()
	content := container.NewVBox()

	if err != nil {
		content = container.NewVBox(container.NewVBox(
			widget.NewLabel(err.Error()),
			widget.NewButton("Home", func() {
				win.Hide()
				h.Show()
				h.SetMaster()
			}),
		))
	} else {
		info, _ := data.GetAll()

		ttl := widget.NewLabel(fmt.Sprintf("Expenses for %v %v", exp.Month, exp.Year))
		budget := widget.NewLabel("")

		if exp.TotalExpenses > info.Budget {
			budget.SetText(fmt.Sprintf("Budget: %v SEK (%v SEK over budget)", info.Budget, exp.TotalExpenses-info.Budget))
		} else {
			budget.SetText(fmt.Sprintf("Budget: %v SEK (%v SEK under budget)", info.Budget, info.Budget-exp.TotalExpenses))
		}

		totalExpenses := widget.NewLabel(fmt.Sprintf("Total Expenses: %v SEK", exp.TotalExpenses))

		groceriesExpenses := widget.NewLabel(fmt.Sprintf("Groceries Expenses: %v SEK (%v", exp.GroceriesExpenses, funcs.Round(100*exp.GroceriesExpenses/exp.TotalExpenses, 2)) + "%)")
		travelExpenses := widget.NewLabel(fmt.Sprintf("Groceries Expenses: %v SEK (%v", exp.TravelExpenses, funcs.Round(100*exp.TravelExpenses/exp.TotalExpenses, 2)) + "%)")
		otherExpenses := widget.NewLabel(fmt.Sprintf("Groceries Expenses: %v SEK (%v", exp.OtherExpenses, funcs.Round(100*exp.OtherExpenses/exp.TotalExpenses, 2)) + "%)")

		content = container.NewVBox(
			ttl,
			budget,
			totalExpenses,
			groceriesExpenses,
			travelExpenses,
			otherExpenses,

			widget.NewButton("Export individual expenses as CSV", func() {
				newWin := a.NewWindow("Export individual expenses")

				fileName := widget.NewEntry()
				fileName.SetText("expenses.csv")

				newWin.SetContent(container.NewVBox(
					fileName,
					widget.NewButton("Save", func() {
						save := dialog.NewFileSave(func(w fyne.URIWriteCloser, e error) {
							if err != nil {
								ErrorMessage(a, newWin, err.Error())
								return
							}

							var out string = "Category,Price\n"

							out += fmt.Sprintf("Total expenses:,%v SEK\n", info.TotalPrice)
							out += fmt.Sprintf("Groceries:,%v SEK\n", info.Groceries.Price)
							out += fmt.Sprintf("Travel:,%v SEK\n", info.Travel.Price)
							out += fmt.Sprintf("Others:,%v SEK\n", info.Others.Price)
							out += fmt.Sprintf("\n")

							out += "Category,Price,Date\n"

							for _, i := range info.Groceries.Items {
								out += fmt.Sprintf("%v,%v SEK,%v-%v-%v\n", i.Category, i.Price, i.Year, i.Month, i.Day)
							}

							for _, i := range info.Travel.Items {
								out += fmt.Sprintf("%v,%v SEK,%v-%v-%v\n", i.Category, i.Price, i.Year, i.Month, i.Day)
							}

							for _, i := range info.Others.Items {
								out += fmt.Sprintf("%v,%v SEK,%v-%v-%v\n", i.Category, i.Price, i.Year, i.Month, i.Day)
							}

							write := []byte(out)
							w.Write(write)
						}, win)

						save.SetFileName(fileName.Text)
						save.Show()
					}),
				))

				newWin.Resize(fyne.NewSize(260, 100))
				newWin.Show()
			}),

			widget.NewButton("Export monthly expenses as CSV", func() {
				newWin := a.NewWindow("Export monthly expenses")

				entry := widget.NewEntry()
				entry.SetText("monthly-expenses")

				newWin.SetContent(container.NewVBox(
					entry,
					widget.NewButton("Save", func() {
						newWin.Hide()
						title := entry.Text + ".csv"

						save := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
							var out string

							out += fmt.Sprintf("Total Expenses:,%v SEK\n", info.TotalPrice)
							out += "\nMonth,Year,Groceries,Travel,Other,Total\n"

							for _, i := range info.Expenses {
								out += fmt.Sprintf("%v,%v,%v,%v,%v,%v\n", i.Month, i.Year, i.GroceriesExpenses, i.TravelExpenses, i.OtherExpenses, i.TotalExpenses)
							}

							w := []byte(out)
							uc.Write(w)
						}, win)

						save.SetFileName(title)
						save.Show()
					}),
				))
				newWin.Resize(fyne.NewSize(260, 100))
				newWin.Show()
			}),

			widget.NewButton("Home", func() {
				h.Show()
				h.SetMaster()
				win.Hide()
			}),
		)
	}

	win.Resize(fyne.NewSize(600, 300))
	win.SetContent(content)
	win.Show()
}

func SetBudget(a fyne.App, h fyne.Window) {
	title := "Set Budget"
	win := a.NewWindow(title)

	info, err := data.GetAll()

	if err != nil {
		ErrorMessage(a, win, err.Error())
		return
	}

	h.Hide()
	win.SetMaster()

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Budget (SEK)")
	entry.SetText(fmt.Sprintf("%v", info.Budget))

	content := container.NewVBox(
		entry,
		widget.NewButton("Set budget", func() {
			budget, err := strconv.ParseFloat(entry.Text, 64)

			if err != nil && budget > 0 {
				ErrorMessage(a, win, "Please only enter numbers")
				return
			}

			if budget <= 0 {
				ErrorMessage(a, win, "Budget cannot be less or equal to zero")
				return
			}

			e := data.UpdateBudget(budget)

			if e == nil {
				ErrorMessage(a, win, "Succesfully set budget to "+fmt.Sprintf("%v", budget))
			}
		}),

		widget.NewButton("Home", func() {
			h.Show()
			h.SetMaster()
			win.Hide()
		}),
	)

	win.Resize(fyne.NewSize(260, 100))
	win.SetContent(content)
	win.Show()
}

func main() {
	a := app.New()
	w := a.NewWindow("Budget Manager")

	title := widget.NewLabel("Budget Manager")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	w.SetContent(container.NewVBox(
		title,
		widget.NewButton("Add an expense", func() { AddExpenses(a, w) }),
		widget.NewButton("See expenses", func() { SeeExpenses(a, w) }),
		widget.NewButton("Set Budget", func() { SetBudget(a, w) }),
	))

	w.Resize(fyne.NewSize(250, 100))
	w.ShowAndRun()
}
