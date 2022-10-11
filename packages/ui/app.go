package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"

	"github.com/alendavid/go-draw/packages/storage"
)

func Run() error {
	app := app.New()
	storage := storage.New()
	window := app.NewWindow("Go Draw")

	functions := Functions(storage)

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Menu", theme.MenuIcon(), Menu(app, storage)),
		container.NewTabItem("Functions", functions),
		container.NewTabItem("Operations", Operations()),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	window.SetContent(tabs)

	// styling options
	window.Resize(fyne.NewSize(960, 720))
	window.CenterOnScreen()

	fmt.Println("Window loaded. Application is ready to run.")

	window.ShowAndRun()

	return nil
}
