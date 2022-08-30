package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Run() error {
	myApp := app.New()
	myWindow := myApp.NewWindow("Go Draw")

	// register other components here

	// styling options
	myWindow.Resize(fyne.NewSize(960, 720))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()

	return nil
}
