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

// returns "A", count of white->black transitions in circuit of neighbors
// of an interior pixel b[r][c]
func transitions(neighbors []int) (count int) {
	bit := neighbors[7]
	for i := 0; i < 8; i++ {
		last := bit
		bit = neighbors[i]

		if last == 0 {
			count += bit
		}
	}
	return count
}

func shouldChangeStep1(x, y int, neighbors []int, img image.Image) bool {
	isBlack := getRedAsOne(img.At(x, y)) == 1
	validSum := (2 <= sum(neighbors)) && (sum(neighbors) <= 8)
	validTransitions := transitions(neighbors) == 1
	c3 := neighbors[0]*neighbors[2]*neighbors[4] == 0 // P2, P4, P6
	c4 := neighbors[2]*neighbors[4]*neighbors[6] == 0 // P4, P6, P8

	return isBlack && validSum && validTransitions && c3 && c4
}

func shouldChangeStep2(x, y int, neighbors []int, img image.Image) bool {
	isBlack := getRedAsOne(img.At(x, y)) == 1
	validSum := (2 <= sum(neighbors)) && (sum(neighbors) <= 8)
	validTransitions := transitions(neighbors) == 1
	c3 := neighbors[0]*neighbors[2]*neighbors[6] == 0 // P2, P4, P8
	c4 := neighbors[0]*neighbors[4]*neighbors[6] == 0 // P2, P6, P8

	return isBlack && validSum && validTransitions && c3 && c4
}

// Zhang-Suen Algorithm
func EdgeThinning(img image.Image) image.Image {
	bounds := img.Bounds()
	m := image.NewRGBA(bounds)

	// fill everything as black
	for i := 1; i < bounds.Dx()-1; i += 1 {
		for j := 1; j < bounds.Dy()-1; j += 1 {
			m.Set(i, j, color.Black)
		}
	}

	var changing [][]int = make([][]int, 0)

	// step 1
	for i := 1; i < bounds.Dx()-1; i += 1 {
		for j := 1; j < bounds.Dy()-1; j += 1 {
			points := neighbors(i, j, img)

			if shouldChangeStep1(i, j, points, img) {
				changing = append(changing, []int{i, j})
			}
		}
	}

	for _, change := range changing {
		x, y := change[0], change[1]

		m.Set(x, y, color.White)
	}

	// step 2
	for i := 1; i < bounds.Dx()-1; i += 1 {
		for j := 1; j < bounds.Dy()-1; j += 1 {
			points := neighbors(i, j, img)

			if shouldChangeStep2(i, j, points, img) {
				changing = append(changing, []int{i, j})
			}
		}
	}

	for _, change := range changing {
		x, y := change[0], change[1]

		m.Set(x, y, color.White)
	}

	return m
}
