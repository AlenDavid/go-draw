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
	tabs := container.NewAppTabs()

	rebuild := func() {}

	var menu, functions, about *container.TabItem

	rebuild = func() {
		fmt.Println("Rebuild is called.")
		tabs.Remove(menu)
		tabs.Remove(functions)
		tabs.Remove(about)

		menu = container.NewTabItemWithIcon("Menu", theme.MenuIcon(), Menu(app, storage, rebuild))
		functions = container.NewTabItem("Functions", Functions(storage, rebuild))
		about = container.NewTabItemWithIcon("About", theme.InfoIcon(), About())

		tabs.Append(menu)
		tabs.Append(functions)
		tabs.Append(about)
	}

	rebuild()

	tabs.SetTabLocation(container.TabLocationTop)

	window.SetContent(tabs)

	// styling options
	window.Resize(fyne.NewSize(960, 720))
	window.CenterOnScreen()

	fmt.Println("Window loaded. Application is ready to run.")

	window.ShowAndRun()

	return nil
}
