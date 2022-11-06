package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/storage"
	ui "github.com/alendavid/go-draw/packages/ui/actions"
)

func Functions(storage storage.Storage, rebuild func()) fyne.CanvasObject {
	box := container.NewAdaptiveGrid(4)

	fmt.Println("box layout", box.Layout)

	_, img := storage.GetWorkingImage()

	if img == nil {
		fmt.Println("Functions is built without image.")
		return widget.NewLabel("Please select an image from Menu tab")
	}

	imageContainer := container.NewWithoutLayout()

	imageCanvas := canvas.NewImageFromImage(img)
	imageCanvas.Resize(fyne.NewSize(float32(img.Bounds().Size().X), float32(img.Bounds().Size().Y)))

	imageContainer.Add(imageCanvas)

	list := container.NewVBox()

	onClick := func() {
		imageContainer.RemoveAll()
		imageContainer.Refresh()
		_, img := storage.GetWorkingImage()

		imageCanvas := canvas.NewImageFromImage(img)
		imageCanvas.Resize(fyne.NewSize(float32(img.Bounds().Size().X), float32(img.Bounds().Size().Y)))
		imageContainer.Resize(fyne.NewSize(float32(img.Bounds().Size().X), float32(img.Bounds().Size().Y)))

		imageContainer.Add(imageCanvas)
		// rebuild()
	}

	list.Add(ui.Translation(storage, onClick))
	list.Add(ui.GrayScale(storage, onClick))
	list.Add(ui.Brightness(storage, onClick))
	list.Add(ui.Contrast(storage, onClick))
	list.Add(ui.Scale(storage, onClick))

	scroll := container.NewVScroll(list)

	box.Add(scroll)
	box.Add(imageContainer)

	fmt.Println("Functions is loaded with image bounds ", img.Bounds())

	return box
}
