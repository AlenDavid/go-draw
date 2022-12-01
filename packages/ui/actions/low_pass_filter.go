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
	thresholdLimit = 3
)

func LowPassFilter(storage storage.Storage, onClick func()) fyne.CanvasObject {
	box := container.NewVBox()

	buttonAverage := widget.NewButton("Apply Low Pass Filter (Average)", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		result := actions.LowPassFilterAverage(workingImage)
		storage.SetImage(result)
		onClick()

		fmt.Println("Done with Apply Low Pass Filter (Average) button.")
	})

	buttonMode := widget.NewButton("Apply Low Pass Filter (Mode)", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		result := actions.LowPassFilterMode(workingImage)
		storage.SetImage(result)
		onClick()

		fmt.Println("Done with Apply Low Pass Filter (Mode) button.")
	})

	box.Add(buttonAverage)
	box.Add(buttonMode)

	return box
}
