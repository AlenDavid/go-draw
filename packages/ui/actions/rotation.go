package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/actions"
	"github.com/alendavid/go-draw/packages/storage"
)

const (
	sliderRotationValue float64 = 90
)

func Rotation(storage storage.Storage, onClick func()) fyne.CanvasObject {
	box := container.NewVBox()
	slider := widget.NewSlider(0, sliderRotationValue)

	button := widget.NewButton("Rotate image", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		result := actions.Rotation(workingImage, slider.Value)
		storage.SetImage(result)
		onClick()

		fmt.Println("Done with rotation")
	})

	box.Add(slider)
	box.Add(button)

	return box
}
