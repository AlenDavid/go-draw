package actions

import (
	"image"
	"image/color"
)

func HighPassRoberts(img image.Image) image.Image {
	bounds := img.Bounds()
	m := image.NewRGBA(bounds)

	for i := 0; i < bounds.Dx(); i += 1 {
		for j := 0; j < bounds.Dy(); j += 1 {

			color := color.RGBA{255, 255, 255, 255}
			m.Set(i, j, color)
		}
	}

	return m
}
