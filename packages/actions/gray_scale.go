package actions

import (
	"image"
	"image/color"
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
