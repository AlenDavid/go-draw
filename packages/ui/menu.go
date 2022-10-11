package ui

import (
	"fmt"
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/storage"
)

func Menu(app fyne.App, storage storage.Storage) fyne.CanvasObject {
	box := container.NewAdaptiveGrid(4)

	imageContainer := container.NewWithoutLayout()

	// imageCanvas := canvas.NewImageFromFile("./assets/lena-255-255.jpeg")
	textCanvas := canvas.NewText("Open an image", color.White)

	imageContainer.Add(textCanvas)

	box.Add(widget.NewButtonWithIcon("Open file", theme.FileImageIcon(), func() {
		window := app.Driver().AllWindows()[0]

		window.Show()

		dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			fmt.Println("Dialog done with interactions.")

			if err != nil || uc == nil {
				fmt.Println(err, uc)
				return
			}

			imageContainer.Remove(imageContainer.Objects[0])

			img2, _, _ := image.Decode(uc)
			storage.SetImage(img2)

			newImageCanvas := canvas.NewImageFromImage(img2)
			newImageCanvas.Resize(fyne.NewSize(float32(img2.Bounds().Size().X), float32(img2.Bounds().Size().Y)))

			imageContainer.Refresh()
			imageContainer.Add(newImageCanvas)

		}, window).Show()
	}))

	box.Add(imageContainer)

	fmt.Println("Menu is loaded.")

	return box
}
