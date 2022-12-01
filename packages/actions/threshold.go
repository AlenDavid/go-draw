package actions

import (
	"image"
	"image/color"
)

func Threshold(value uint8, img image.Image) image.Image {
	bounds := img.Bounds()
	m := image.NewGray(bounds)

	for i := 0; i < bounds.Dx(); i += 1 {
		for j := 0; j < bounds.Dy(); j += 1 {
			R, G, B, _ := img.At(i, j).RGBA()
			Y := (0.299*float64(R) + 0.587*float64(G) + 0.114*float64(B)) / 256

			if Y > float64(value) {
				m.Set(i, j, color.White)
			} else {
				m.Set(i, j, color.Black)
			}
		}
	}

	return m
}
