package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/actions"
	"github.com/alendavid/go-draw/packages/storage"
)

func Mirroring(storage storage.Storage, onClick func()) fyne.CanvasObject {
	box := container.NewVBox()
	var direction actions.MirrorDirection = actions.HORIZONTAL

	radio := widget.NewRadioGroup([]string{"horizontal", "vertical"}, func(s string) {
		direction = actions.MirrorDirection(s)
	})

	button := widget.NewButton("Apply mirror", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		result := actions.Mirroring(workingImage, direction)

		storage.SetImage(result)
		onClick()

		fmt.Println("Done with Mirroring button.")
	})

	box.Add(radio)
	box.Add(button)

	return box
}
