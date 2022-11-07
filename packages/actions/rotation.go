package actions

import (
	"image"
	"image/color"
	"math"
)

// 0 >= degree < 360
// img = [100, 100]
func Rotation(img image.Image, degree float64) image.Image {
	var piDegree float64 = math.Pi * (30.0 / 180.0)
	var cos, sin float64 = math.Cos(piDegree), math.Sin(piDegree)
	size := img.Bounds().Size()

	m := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))

	var xd, yd int
	var xc, yc int = size.X / 2, size.Y / 2

	for x0 := 0; x0 < size.X; x0++ { // 0, 100
		for y0 := 0; y0 < size.Y; y0++ { // 0, 100
			xd = int(cos*float64(x0-xc)+sin*float64(y0-yc)) + xc
			yd = int((-sin)*float64(x0-xc)+cos*float64(y0-yc)) + yc

			if !image.Pt(xd, yd).In(img.Bounds()) {
				m.Set(x0, y0, color.Black)
			} else {
				m.Set(x0, y0, img.At(xd, yd))
			}
		}
	}

	return m
}
