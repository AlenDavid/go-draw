package ui

import (
	"fmt"
	"image"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/alendavid/go-draw/packages/actions"
)

func Run() error {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")

	myContainer := container.NewAdaptiveGrid(2)
	myCanvas := myWindow.Canvas()

	path := "./assets/lena-255-255.jpeg"
	lena, err := os.Open(path)
	defer lena.Close()

	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(lena)

	if err != nil {
		fmt.Println(err)
	}

	imageContainer := container.NewWithoutLayout()
	imageCanvas := canvas.NewImageFromImage(actions.GrayScale(img))
	imageCanvas.Resize(fyne.NewSize(255, 255))

	imageContainer.Add(imageCanvas)

	button := widget.NewButton("click me", func() {
		fmt.Println("click")
		imageContainer.Remove(imageContainer.Objects[0])

		dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			if err != nil || uc == nil {
				fmt.Println(err, uc)
				return
			}

			img2, _, _ := image.Decode(uc)
			newImageCanvas := canvas.NewImageFromImage(actions.GrayScale(img2))
			newImageCanvas.Resize(fyne.NewSize(255, 255))

			imageContainer.Add(newImageCanvas)
		}, myWindow).Show()
	})

	myContainer.Add(imageContainer)
	myContainer.Add(button)
	myCanvas.SetContent(myContainer)

	myWindow.Resize(fyne.NewSize(500, 500))
	myWindow.ShowAndRun()

	return nil
}
