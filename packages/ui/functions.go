package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/storage"
)

func Functions(storage storage.Storage) fyne.CanvasObject {
	box := container.NewAdaptiveGrid(4)

	err, img := storage.GetWorkingImage()

	if err == nil || img == nil {
		fmt.Println("Functions is without image.")
		return widget.NewLabel("Please select an image from Menu tab")
	}

	imageContainer := container.NewWithoutLayout()

	imageCanvas := canvas.NewImageFromImage(img)
	imageCanvas.Resize(fyne.NewSize(float32(img.Bounds().Size().X), float32(img.Bounds().Size().Y)))

	imageContainer.Add(imageCanvas)
	box.Add(imageContainer)

	fmt.Println("Functions is loaded.")

	return box
}
