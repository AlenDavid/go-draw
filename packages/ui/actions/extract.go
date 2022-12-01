package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/alendavid/go-draw/packages/actions"
	"github.com/alendavid/go-draw/packages/storage"
)

func ExtractCharacteristics(storage storage.Storage, onClick func()) fyne.CanvasObject {
	box := container.NewVBox()

	results := widget.NewTextGrid()

	results.SetText("Please use Extract Characteristics first.")

	buttonAverage := widget.NewButton("Extract Characteristics", func() {
		err, workingImage := storage.GetWorkingImage()

		if err != nil {
			panic(err)
		}

		text := ""

		for _, result := range actions.ExtractCharacteristics(workingImage) {
			text += fmt.Sprintln(result)
		}

		results.SetText(text)

		onClick()

		fmt.Println("Done with Extract Characteristics button.")
	})

	box.Add(buttonAverage)
	box.Add(results)

	return box
}
