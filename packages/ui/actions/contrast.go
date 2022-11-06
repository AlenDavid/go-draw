package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/actions"
	"github.com/alendavid/go-draw/packages/storage"
)

func Contrast(storage storage.Storage, onClick func()) fyne.CanvasObject {
	return widget.NewButton("Apply 10% contrast", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		result := actions.Contrast(workingImage, 1.1)
		storage.SetImage(result)
		onClick()
		fmt.Println("Done with Contrast button.")
	})
}
