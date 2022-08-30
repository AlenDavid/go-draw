package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Run() error {
	app := app.New()
	window := app.NewWindow("Go Draw")

	// register other components here

	// styling options
	window.Resize(fyne.NewSize(960, 720))
	window.CenterOnScreen()
	window.ShowAndRun()

	return nil
}
