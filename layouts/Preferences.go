package layouts

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func setBudget(root fyne.App, home fyne.Window) {
	window := root.NewWindow("Set budget - BudgeTea")

	budgetEntry := widget.NewEntry()
	budgetEntry.SetPlaceHolder("New Budget")

	window.SetContent(container.NewVBox(
		budgetEntry,

		widget.NewButton("Set budget", func() {

		}),
	))

	window.Resize(fyne.NewSize(100, 100))
	window.Show()
}

func setDenomination(root fyne.App, home fyne.Window) {
	var denomination string

	window := root.NewWindow("Set denomination - BudgeTea")

	title := widget.NewLabel("Denominations: ")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	denoms := widget.NewRadioGroup([]string{
		"EUR",
		"SEK",
		"USD",
	}, func(choice string) {
		denomination = choice
	})

	window.SetContent(container.NewVBox(
		denoms,

		widget.NewButton("Set denomination", func() {
			fmt.Println(denomination)
		}),
	))
}

func Preferences(root fyne.App, home fyne.Window) {
	window := root.NewWindow("Preferences - BudgeTea")
	window.SetMaster()
	home.Hide()

	title := widget.NewLabel("Preferences")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	window.SetContent(container.NewVBox(
		title,

		widget.NewButton("Set Budget", func() {
			setBudget(root, window)
		}),
		widget.NewButton("Change preferred denomination", func() {
			setDenomination(root, window)
		}),
		widget.NewButton("Home", func() {
			home.SetMaster()
			home.Show()
			window.Hide()
		}),
	))

	window.Resize(fyne.NewSize(300, 100))
	window.Show()
}
