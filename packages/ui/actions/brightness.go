package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/actions"
	"github.com/alendavid/go-draw/packages/storage"
)

func Brightness(storage storage.Storage, onClick func()) fyne.CanvasObject {
	return widget.NewButton("Apply 10 brightness", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		result := actions.Brightness(workingImage, 10)
		storage.SetImage(result)
		onClick()

		fmt.Println("Done with GrayScale button.")
	})
}
