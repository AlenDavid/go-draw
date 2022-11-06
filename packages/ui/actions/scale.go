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
	sliderScaleValue float64 = 200
)

func Scale(storage storage.Storage, onClick func()) fyne.CanvasObject {
	box := container.NewVBox()
	sliderX := widget.NewSlider(1, sliderScaleValue)
	sliderY := widget.NewSlider(1, sliderScaleValue)

	button := widget.NewButton("Apply scale", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		x, y := sliderX.Value/100, sliderY.Value/100

		result := actions.Scale(workingImage, x, y)
		storage.SetImage(result)
		onClick()

		fmt.Println("Done with Translation button.")
	})

	box.Add(sliderX)
	box.Add(sliderY)
	box.Add(button)

	return box
}
