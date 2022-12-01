package actions

import (
	"image"
	"image/color"
	"math"
)

const (
	HighPassRobertsThreshold = 100
)

func HighPassRoberts(img image.Image) image.Image {
	bounds := img.Bounds()
	m := image.NewRGBA(bounds)

	xKernel := [][]int{{0, 0, 0}, {0, -1, 0}, {0, 0, 1}}
	yKernel := [][]int{{0, 0, 0}, {0, 0, -1}, {0, 1, 0}}

	for i := 1; i < bounds.Dx()-1; i += 1 {
		for j := 1; j < bounds.Dy()-1; j += 1 {
			var r uint8
			var gx, gy int = 0, 0

			for k := 0; k < 3; k += 1 {
				for l := 0; l < 3; l += 1 {
					color := img.At(i+(k-1), j+(l-1))
					R, _, _, _ := color.RGBA()

					gx += int(R/257) * xKernel[k][l]
					gy += int(R/257) * yKernel[k][l]
				}
			}

			if math.Sqrt(math.Pow(float64(gx), 2)+math.Pow(float64(gy), 2)) > HighPassRobertsThreshold {
				r = 255
			}

			color := color.RGBA{r, r, r, 255}
			m.Set(i, j, color)
		}
	}

	return m
}

func HighPassSobel(img image.Image) image.Image {
	bounds := img.Bounds()
	m := image.NewRGBA(bounds)

	xKernel := [][]int{{1, 0, -1}, {2, 0, -2}, {1, 0, -1}}
	yKernel := [][]int{{1, 2, 1}, {0, 0, 0}, {-1, -2, -1}}

	for i := 1; i < bounds.Dx()-1; i += 1 {
		for j := 1; j < bounds.Dy()-1; j += 1 {
			var r uint8
			var gx, gy int = 0, 0

			for k := 0; k < 3; k += 1 {
				for l := 0; l < 3; l += 1 {
					color := img.At(i+(k-1), j+(l-1))
					R, _, _, _ := color.RGBA()

					gx += int(R/257) * xKernel[k][l]
					gy += int(R/257) * yKernel[k][l]
				}
			}

			if math.Sqrt(math.Pow(float64(gx), 2)+math.Pow(float64(gy), 2)) > HighPassRobertsThreshold {
				r = 255
			}

			color := color.RGBA{r, r, r, 255}
			m.Set(i, j, color)
		}
	}

	return m
}
