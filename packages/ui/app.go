package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func Run() error {
	app := app.New()
	window := app.NewWindow("Go Draw")

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Menu", theme.MenuIcon(), Menu(app)),
		container.NewTabItem("Functions", Functions()),
		container.NewTabItem("Operations", Operations()),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	window.SetContent(tabs)

	// styling options
	window.Resize(fyne.NewSize(960, 720))
	window.CenterOnScreen()
	window.ShowAndRun()

	return nil
}
