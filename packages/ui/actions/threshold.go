package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/actions"
	"github.com/alendavid/go-draw/packages/storage"
)

func Threshold(storage storage.Storage, onClick func()) fyne.CanvasObject {
	box := container.NewVBox()
	sliderX := widget.NewSlider(1, 254)

	button1 := widget.NewButton("Apply Threshold", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		var value uint8 = uint8(sliderX.Value)

		result := actions.Threshold(value, workingImage)
		storage.SetImage(result)
		onClick()

		fmt.Println("Done with Apply Threshold button.")
	})

	box.Add(sliderX)
	box.Add(button1)

	return box
}
