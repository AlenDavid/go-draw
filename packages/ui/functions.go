package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/actions"
	"github.com/alendavid/go-draw/packages/storage"
	ui "github.com/alendavid/go-draw/packages/ui/actions"
)

func Functions(storage storage.Storage) fyne.CanvasObject {
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

	list.Add(ui.GrayScale(storage, func() {
		box.Remove(imageCanvas)
		_, img := storage.GetWorkingImage()

		imageCanvas := canvas.NewImageFromImage(actions.GrayScale(img))
		imageCanvas.Resize(fyne.NewSize(float32(img.Bounds().Size().X), float32(img.Bounds().Size().Y)))

		imageContainer.Add(imageCanvas)
	}))

	list.Add(ui.Brightness(storage, func() {
		box.Remove(imageCanvas)
		_, img := storage.GetWorkingImage()

		imageCanvas := canvas.NewImageFromImage(actions.GrayScale(img))
		imageCanvas.Resize(fyne.NewSize(float32(img.Bounds().Size().X), float32(img.Bounds().Size().Y)))

		imageContainer.Add(imageCanvas)
	}))

	scroll := container.NewVScroll(list)

	box.Add(scroll)

	box.Add(imageContainer)

	fmt.Println("Functions is loaded with image bounds ", img.Bounds())

	return box
}
