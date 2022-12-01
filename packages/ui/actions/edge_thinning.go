package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/actions"
	"github.com/alendavid/go-draw/packages/storage"
)

func EdgeThinning(storage storage.Storage, onClick func()) fyne.CanvasObject {
	box := container.NewVBox()

	buttonAverage := widget.NewButton("Apply Edge Thinning", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		result := actions.EdgeThinning(workingImage)
		storage.SetImage(result)
		onClick()

		fmt.Println("Done with Apply Edge Thinning button.")
	})

	box.Add(buttonAverage)

	return box
}
