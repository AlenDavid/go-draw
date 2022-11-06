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
	sliderValue = 10
)

func Translation(storage storage.Storage, onClick func()) fyne.CanvasObject {
	box := container.NewVBox()
	sliderX := widget.NewSlider(-sliderValue, sliderValue)
	sliderY := widget.NewSlider(-sliderValue, sliderValue)

	button := widget.NewButton("Apply translation", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		x, y := int(sliderX.Value), int(sliderY.Value)

		result := actions.Translation(workingImage, x, y)
		storage.SetImage(result)
		onClick()

		fmt.Println("Done with Translation button.")
	})

	box.Add(sliderX)
	box.Add(sliderY)
	box.Add(button)

	return box
}
