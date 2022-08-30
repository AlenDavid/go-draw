package ui

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func GrayScale(img image.Image) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, img.Bounds().Size().X, img.Bounds().Size().Y))

	for i := 0; i < img.Bounds().Size().X; i += 1 {
		for j := 0; j < img.Bounds().Size().Y; j += 1 {
			r, g, b, _ := img.At(i, j).RGBA() // 0 - 65535

			var value uint8 = uint8(((r + g + b) * 255) / (65535 * 3))
			c := color.Gray{value} // 0 - 255

			m.Set(i, j, c)
		}
	}

	return m
}

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
	imageCanvas := canvas.NewImageFromImage(GrayScale(img))
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
			newImageCanvas := canvas.NewImageFromImage(GrayScale(img2))
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
