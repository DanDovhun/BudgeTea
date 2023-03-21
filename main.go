package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func main() {
	root := app.New()
	home := root.NewWindow("Budgetie")

	label := widget.NewLabel("Budgetie")
	label.Position()

	home.SetContent(container.NewVBox(
		label,
	))
	home.Resize(fyne.NewSize(400, 200))
	home.ShowAndRun()
}
