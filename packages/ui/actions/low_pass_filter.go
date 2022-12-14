package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/actions"
	"github.com/alendavid/go-draw/packages/storage"
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

	buttonMedian := widget.NewButton("Apply Low Pass Filter (Median)", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		result := actions.LowPassFilterMedian(workingImage)
		storage.SetImage(result)
		onClick()

		fmt.Println("Done with Apply Low Pass Filter (Median) button.")
	})

	box.Add(buttonAverage)
	box.Add(buttonMode)
	box.Add(buttonMedian)

	return box
}
