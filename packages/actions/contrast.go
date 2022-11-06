package actions

import (
	"image"
	"image/color"
)

// a is our current color value
// c is the contrast factor
func calcContrast(a uint8, c float32) uint8 {
	if c == 1 {
		return a
	}

	fa := float32(a) * c

	if fa > 255 {
		return 255
	}

	return uint8(fa)
}

func Contrast(img image.Image, C float32) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, img.Bounds().Size().X, img.Bounds().Size().Y))

	for i := 0; i < img.Bounds().Size().X; i += 1 {
		for j := 0; j < img.Bounds().Size().Y; j += 1 {
			// 0 - 4,294,967,295, but no value will be higher than 65,535
			// this behavior allows R | G | B to be multiplied for A without overflow
			r, g, b, a := img.At(i, j).RGBA()

			// by dividing with 257 and transforming to uint8, we get the expected 0-255 value we need
			var r1, g1, b1 uint8 = calcContrast(uint8(r/257), C), calcContrast(uint8(g/257), C), calcContrast(uint8(b/257), C)

			c := color.RGBA{r1, g1, b1, uint8(a)}

			m.Set(i, j, c)
		}
	}

	return m
}
