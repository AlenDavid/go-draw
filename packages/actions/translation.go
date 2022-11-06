package actions

import (
	"image"
	"image/color"
)

func isOutOfBounds(img image.Image, x, y int) bool {
	if x < 0 || y < 0 {
		return true
	}

	return x >= img.Bounds().Size().X || y >= img.Bounds().Size().Y
}

func Translation(img image.Image, translationX, translationY int) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, img.Bounds().Size().X, img.Bounds().Size().Y))

	for i := 0; i < img.Bounds().Size().X; i += 1 {
		for j := 0; j < img.Bounds().Size().Y; j += 1 {
			// if translation is out of bounds, set white and continue
			if isOutOfBounds(img, i-translationX, j-translationY) {
				m.Set(i, j, color.White)
				continue
			}

			c := img.At(i-translationX, j-translationY)

			m.Set(i, j, c)
		}
	}

	return m
}
