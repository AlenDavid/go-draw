package actions

import (
	"fmt"
	"image"
	"image/color"
)

// get this pixels from around the image
// c = (x, y)
// _, 1, _
// 1, c, 1
// _, 1, _
func minCross(img image.Image, x, y int, factor int) (c color.Color) {
	var r, g, b, a int = 255, 255, 255, 255
	var pixesAt [][]int = make([][]int, 0)

	pixesAt = append(pixesAt,
		[]int{x, y - 1},
		[]int{x - 1, y}, []int{x, y}, []int{x + 1, y},
		[]int{x, y + 1},
	)

	for _, p := range pixesAt {
		fr, fg, fb, fa := img.At(p[0], p[1]).RGBA()
		fri, fgi, fbi, fai := int(fr/257.0), int(fg/257.0), int(fb/257.0), int(fa/257.0)

		if r > fri {
			r = fri
		}

		if g > fgi {
			g = fgi
		}

		if b > fbi {
			b = fbi
		}

		if a > fai {
			a = fai
		}
	}

	return color.RGBA{
		uint8MinMax(r - factor),
		uint8MinMax(g - factor),
		uint8MinMax(b - factor),
		uint8MinMax(a),
	}
}

// get this pixels from around the image
// c = (x, y)
// 1, 1, 1
// 1, c, 1
// 1, 1, 1
func minRound(img image.Image, x, y, factor int) (c color.Color) {
	var r, g, b, a int = 255, 255, 255, 255
	var pixesAt [][]int = make([][]int, 0)

	pixesAt = append(pixesAt,
		[]int{x - 1, y - 1}, []int{x, y - 1}, []int{x + 1, y - 1},
		[]int{x - 1, y}, []int{x, y}, []int{x + 1, y},
		[]int{x - 1, y + 1}, []int{x, y + 1}, []int{x + 1, y + 1},
	)

	for _, p := range pixesAt {
		fr, fg, fb, fa := img.At(p[0], p[1]).RGBA()
		fri, fgi, fbi, fai := int(fr/257.0), int(fg/257.0), int(fb/257.0), int(fa/257.0)

		if r > fri {
			r = fri
		}

		if g > fgi {
			g = fgi
		}

		if b > fbi {
			b = fbi
		}

		if a > fai {
			a = fai
		}
	}

	return color.RGBA{
		uint8MinMax(r - factor),
		uint8MinMax(g - factor),
		uint8MinMax(b - factor),
		uint8MinMax(a),
	}
}

// get this pixels from around the image
// c = (x, y)
// 1, _, 1
// _, c, _
// 1, _, 1
func minDiagonals(img image.Image, x, y, factor int) (c color.Color) {
	var r, g, b, a int = 255, 255, 255, 255
	var pixesAt [][]int = make([][]int, 0)

	pixesAt = append(pixesAt,
		[]int{x - 1, y - 1}, []int{x + 1, y - 1},
		[]int{x, y},
		[]int{x - 1, y + 1}, []int{x + 1, y + 1},
	)

	for _, p := range pixesAt {
		fr, fg, fb, fa := img.At(p[0], p[1]).RGBA()
		fri, fgi, fbi, fai := int(fr/257.0), int(fg/257.0), int(fb/257.0), int(fa/257.0)

		if r > fri {
			r = fri
		}

		if g > fgi {
			g = fgi
		}

		if b > fbi {
			b = fbi
		}

		if a > fai {
			a = fai
		}
	}

	return color.RGBA{
		uint8MinMax(r - factor),
		uint8MinMax(g - factor),
		uint8MinMax(b - factor),
		uint8MinMax(a),
	}
}

// 0 > factor <= 2
func Erosion(img image.Image, dt DilatationType, factor int) image.Image {
	switch dt {
	case CROSS:
		{
			return applyStructure(img, factor, minCross)
		}
	case ROUND:
		{
			return applyStructure(img, factor, minRound)
		}
	case DIAGONALS:
		{
			return applyStructure(img, factor, minDiagonals)
		}

	default:
		fmt.Println("got some wrong type")
	}

	return img
}
