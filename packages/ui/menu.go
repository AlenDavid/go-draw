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
)

func Menu(app fyne.App) fyne.CanvasObject {
	box := container.NewAdaptiveGrid(4)

	imageContainer := container.NewWithoutLayout()
	imageCanvas := canvas.NewText("Open an image first", color.White)

	imageContainer.Add(imageCanvas)

	box.Add(widget.NewButtonWithIcon("Open file", theme.FileImageIcon(), func() {
		window := app.Driver().AllWindows()[0]

		window.Show()

		dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			fmt.Println("dialog done with interactions")
			if err != nil || uc == nil {
				fmt.Println(err, uc)
				return
			}

			imageContainer.Remove(imageContainer.Objects[0])

			img2, _, _ := image.Decode(uc)
			newImageCanvas := canvas.NewImageFromImage(img2)
			newImageCanvas.Resize(fyne.NewSize(255, 255))

			imageContainer.Add(newImageCanvas)
			imageContainer.Refresh()

		}, window).Show()
	}))

	box.Add(imageCanvas)

	return box
}
