package actions

import (
	"fmt"
	"image"
	"image/color"
)

type DilatationType string

const (
	CROSS     DilatationType = "cross"
	ROUND     DilatationType = "round"
	DIAGONALS DilatationType = "diagonals"
)

func uint8MinMax(i int) uint8 {
	if i > 255 {
		return 255
	}

	if i < 0 {
		return 0
	}

	return uint8(i)
}

// get this pixels from around the image
// c = (x, y)
// _, 1, _
// 1, c, 1
// _, 1, _
func maxCross(img image.Image, x, y int, factor int) (c color.Color) {
	_, _, _, _ = img.At(x, y).RGBA()

	var r, g, b, a int
	var pixesAt [][]int = make([][]int, 0)

	pixesAt = append(pixesAt,
		[]int{x, y - 1},
		[]int{x - 1, y}, []int{x, y}, []int{x + 1, y},
		[]int{x, y + 1},
	)

	for _, p := range pixesAt {
		fr, fg, fb, fa := img.At(p[0], p[1]).RGBA()
		fri, fgi, fbi, fai := int(fr/257.0), int(fg/257.0), int(fb/257.0), int(fa/257.0)

		if r < fri {
			r = fri
		}

		if g < fgi {
			g = fgi
		}

		if b < fbi {
			b = fbi
		}

		if a < fai {
			a = fai
		}
	}

	return color.RGBA{
		uint8MinMax(r + factor),
		uint8MinMax(g + factor),
		uint8MinMax(b + factor),
		uint8MinMax(a + factor),
	}
}

func applyCross(img image.Image, factor int) image.Image {
	size := img.Bounds().Size()
	m := image.NewRGBA(img.Bounds())

	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			m.Set(i, j, maxCross(img, i, j, factor))
		}
	}

	return m
}

// 0 > factor <= 2
func Dilatation(img image.Image, dt DilatationType, factor int) image.Image {
	switch dt {
	case CROSS:
		{
			return applyCross(img, factor)
		}
	case ROUND:
		{
			fmt.Println("not implemented")
		}
	case DIAGONALS:
		{
			fmt.Println("not implemented")
		}

	default:
		fmt.Println("got some wrong type")
	}

	return img
}
