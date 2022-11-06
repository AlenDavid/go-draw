package actions

import (
	"fmt"
	"image"
)

// 0 > factor <= 2
func Scale(img image.Image, factorX, factorY float64) image.Image {
	size := img.Bounds().Size()
	x1, y1 := int(float64(size.X)*factorX), int(float64(size.Y)*factorY)
	fmt.Printf("Scaling x=%d y=%d per factors %f %f with values %d %d\n", img.Bounds().Size().X, img.Bounds().Size().Y, factorX, factorY, x1, y1)

	m := image.NewRGBA(image.Rect(0, 0, x1, y1))

	for i := 0; i < x1; i += 1 {
		for j := 0; j < y1; j += 1 {
			// if factor = 2
			// [final matrix] = [base matrix]
			// [0, 0] = [0, 0]
			// [1, 1] = [0.5, 0.5]
			// [2, 2] = [1, 1]
			ix, iy := int(float64(i)/factorX), int(float64(j)/factorY)
			c := img.At(ix, iy)

			m.Set(i, j, c)
		}
	}

	return m
}
