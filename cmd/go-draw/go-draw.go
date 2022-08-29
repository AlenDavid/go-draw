package main

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
)

type Drawing struct {
	image.Image
}

func GrayScale(img image.Image) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, img.Bounds().Size().X, img.Bounds().Size().Y))

	for i := 0; i < img.Bounds().Size().X; i += 1 {
		for j := 0; j < img.Bounds().Size().Y; j += 1 {
			var c color.Color = img.At(i, j)
			m.Set(i, j, c)
		}
	}

	return m
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")

	myContainer := container.NewWithoutLayout()
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

	imageCanva := canvas.NewImageFromImage(GrayScale(img))
	imageCanva.Resize(fyne.NewSize(255, 255))

	myContainer.Add(imageCanva)
	myCanvas.SetContent(myContainer)

	myWindow.Resize(fyne.NewSize(500, 500))
	myWindow.ShowAndRun()
}
