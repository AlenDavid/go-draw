package buttons

import (
	"fmt"
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/storage"
)

func OpenFile(app fyne.App, container *fyne.Container, storage storage.Storage, rebuild func()) *widget.Button {
	return widget.NewButtonWithIcon("Open file", theme.FileImageIcon(), func() {
		window := app.Driver().AllWindows()[0]

		window.Show()

		dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			fmt.Println("Dialog done with interactions.")

			if err != nil || uc == nil {
				fmt.Println(err, uc)
				return
			}

			container.Remove(container.Objects[0])

			img2, _, _ := image.Decode(uc)
			storage.SetImage(img2)

			newImageCanvas := canvas.NewImageFromImage(img2)
			newImageCanvas.Resize(fyne.NewSize(float32(img2.Bounds().Size().X), float32(img2.Bounds().Size().Y)))

			container.Refresh()
			rebuild()
			container.Add(newImageCanvas)

		}, window).Show()
	})
}

// TODO: work on this function
func SaveFile(app fyne.App, container *fyne.Container, storage storage.Storage, rebuild func()) *widget.Button {
	return widget.NewButtonWithIcon("Save file", theme.FileImageIcon(), func() {
		window := app.Driver().AllWindows()[0]

		window.Show()

		dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			fmt.Println("Dialog done with interactions.")

			if err != nil || uc == nil {
				fmt.Println(err, uc)
				return
			}

			container.Remove(container.Objects[0])

			img2, _, _ := image.Decode(uc)
			storage.SetImage(img2)

			newImageCanvas := canvas.NewImageFromImage(img2)
			newImageCanvas.Resize(fyne.NewSize(float32(img2.Bounds().Size().X), float32(img2.Bounds().Size().Y)))

			container.Refresh()
			rebuild()
			container.Add(newImageCanvas)

		}, window).Show()
	})
}
