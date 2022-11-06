package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/alendavid/go-draw/packages/storage"
	"github.com/alendavid/go-draw/packages/ui/buttons"
)

func Menu(app fyne.App, storage storage.Storage, rebuild func()) fyne.CanvasObject {
	box := container.NewAdaptiveGrid(4)

	imageContainer := container.NewWithoutLayout()

	textCanvas := canvas.NewText("Open an image", color.White)

	imageContainer.Add(textCanvas)

	openFileButton := buttons.OpenFile(app, imageContainer, storage, rebuild)
	saveFileButton := buttons.SaveFile(app, imageContainer, storage, rebuild)

	list := container.NewVBox()

	list.Add(openFileButton)
	list.Add(saveFileButton)

	box.Add(list)

	box.Add(imageContainer)

	fmt.Println("Menu is loaded.")

	return box
}
