package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/actions"
	"github.com/alendavid/go-draw/packages/storage"
)

func HighPassFilter(storage storage.Storage, onClick func()) fyne.CanvasObject {
	box := container.NewVBox()

	button1 := widget.NewButton("Apply High Pass Filter (1)", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		result := actions.HighPassRoberts(workingImage)
		storage.SetImage(result)
		onClick()

		fmt.Println("Done with Apply High Pass Filter (1) button.")
	})

	button2 := widget.NewButton("Apply High Pass Filter (2)", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		result := actions.HighPassRoberts(workingImage)
		storage.SetImage(result)
		onClick()

		fmt.Println("Done with Apply High Pass Filter (2) button.")
	})

	button3 := widget.NewButton("Apply High Pass Filter (3)", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		result := actions.HighPassRoberts(workingImage)
		storage.SetImage(result)
		onClick()

		fmt.Println("Done with Apply High Pass Filter (3) button.")
	})

	box.Add(button1)
	box.Add(button2)
	box.Add(button3)

	return box
}
