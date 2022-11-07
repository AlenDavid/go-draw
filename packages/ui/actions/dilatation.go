package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/actions"
	"github.com/alendavid/go-draw/packages/storage"
)

func Dilatation(storage storage.Storage, onClick func()) fyne.CanvasObject {
	box := container.NewVBox()
	var direction actions.DilatationType = actions.CROSS

	radioType := widget.NewRadioGroup([]string{"cross", "round", "diagonals"}, func(s string) {
		direction = actions.DilatationType(s)
	})

	radioType.Selected = string(direction)

	button := widget.NewButton("Apply dilatation", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		result := actions.Dilatation(workingImage, direction, 10)

		storage.SetImage(result)
		onClick()

		fmt.Println("Done with Dilatation button.")
	})

	box.Add(radioType)
	box.Add(button)

	return box
}
