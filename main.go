package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Clinic App")

	label := widget.NewLabel("Clinic system booted...")
	w.SetContent(label)

	w.ShowAndRun()
}
