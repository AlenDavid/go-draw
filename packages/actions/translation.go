package actions

import (
	"image"
	"image/color"
)

func Translation(img image.Image, translationX, translationY int) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, img.Bounds().Size().X, img.Bounds().Size().Y))

	for i := 0; i < img.Bounds().Size().X; i += 1 {
		for j := 0; j < img.Bounds().Size().Y; j += 1 {
			// if translation is out of bounds, set black and continue
			if !image.Pt(i-translationX, j-translationY).In(img.Bounds()) {
				m.Set(i, j, color.Black)
				continue
			}

			c := img.At(i-translationX, j-translationY)

			m.Set(i, j, c)
		}
	}

	return m
}
