package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/actions"
	"github.com/alendavid/go-draw/packages/storage"
)

func GrayScale(storage storage.Storage, onClick func()) fyne.CanvasObject {
	return widget.NewButton("Apply gray scale", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		result := actions.GrayScale(workingImage)
		storage.SetImage(result)
		onClick()

		fmt.Println("Done with GrayScale button.")
	})
}
