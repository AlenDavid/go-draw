package actions

import (
	"image"
	"image/color"
)

func getRedAsOne(c color.Color) int {
	R, _, _, _ := c.RGBA()

	if int(R/257) == 0 {
		return 0
	}

	return 1
}

func neighbors(x, y int, img image.Image) []int {
	x_1, y_1, x1, y1 := x-1, y-1, x+1, y+1

	return []int{
		getRedAsOne(img.At(x_1, y)),   // P2
		getRedAsOne(img.At(x_1, y1)),  // P3
		getRedAsOne(img.At(x, y1)),    // P4
		getRedAsOne(img.At(x1, y1)),   // P5
		getRedAsOne(img.At(x1, y)),    // P6
		getRedAsOne(img.At(x_1, y_1)), // P7
		getRedAsOne(img.At(x, y_1)),   // P8
		getRedAsOne(img.At(x_1, y_1)), // P9
	}
}

func sum(arr []int) (res int) {
	for _, v := range arr {
		res += v
	}

	return
}

func shouldChange(x, y int, neighbors []int, img image.Image) bool {
	return true
}

// Zhang-Suen Algorithm
func EdgeThinning(img image.Image) image.Image {
	bounds := img.Bounds()
	m := image.NewRGBA(bounds)

	var changing [][]int = make([][]int, 0)

	// step 1
	for i := 1; i < bounds.Dx()-1; i += 1 {
		for j := 1; j < bounds.Dy()-1; j += 1 {
			points := neighbors(i, j, img)

			if shouldChange(i, j, points, img) {
				changing = append(changing, []int{i, j})
			}
		}
	}

	// step 2

	for _, change := range changing {
		x, y := change[0], change[1]

		m.Set(x, y, color.White)
	}

	return m
}
