package actions

import (
	"image"
	"image/color"
)

func Brightness(img image.Image, B int) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, img.Bounds().Size().X, img.Bounds().Size().Y))

	for i := 0; i < img.Bounds().Size().X; i += 1 {
		for j := 0; j < img.Bounds().Size().Y; j += 1 {
			r, g, b, a := img.At(i, j).RGBA() // 0 - 65535

			c := color.RGBA{uint8(int(r) + B), uint8(int(g) + B), uint8(int(b) + B), uint8(a)} // 0 - 255

			m.Set(i, j, c)
		}
	}

	return m
}
