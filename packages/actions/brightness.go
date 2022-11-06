package actions

import (
	"image"
	"image/color"
)

func calcBrightness(a uint8, b int8) uint8 {
	// for a = 250 and b = 10, we must return 255
	// for a = 50 and b = 10, we must return 60
	if b >= 0 {
		diff := 255 - a // 5, 205
		if diff < uint8(b) {
			// since 5 is less than 10, a + b = 250 which would overflow a
			return 255
		}

		// otherwise, we can safely return 60
		return a + uint8(b)
	}

	// for a = 10 and b = -20, we must return 0
	// for a = 200 and b = -20, we must return 180
	diff := uint8(-b) // 20
	if b < 0 {
		// case 10 < 20, return 0
		if a < diff {
			return 0
		}
	}

	// default case:
	// we know a is bigger than the positive number of brightness so we can safely assume
	// we can subtract the positive value of b from a
	return a - diff
}

func Brightness(img image.Image, B int8) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, img.Bounds().Size().X, img.Bounds().Size().Y))

	for i := 0; i < img.Bounds().Size().X; i += 1 {
		for j := 0; j < img.Bounds().Size().Y; j += 1 {
			// 0 - 4,294,967,295, but no value will be higher than 65,535
			// this behavior allows R | G | B to be multiplied for A without overflow
			r, g, b, a := img.At(i, j).RGBA()

			// by dividing with 257 and transforming to uint8, we get the expected 0-255 value we need
			var r1, g1, b1 uint8 = calcBrightness(uint8(r/257), B),
				calcBrightness(uint8(g/257), B),
				calcBrightness(uint8(b/257), B)

			c := color.RGBA{r1, g1, b1, uint8(a)}

			m.Set(i, j, c)
		}
	}

	return m
}
