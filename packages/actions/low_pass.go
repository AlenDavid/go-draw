package actions

import (
	"image"
	"image/color"
	"sort"
)

func LowPassFilterAverage(img image.Image) image.Image {
	bounds := img.Bounds()
	m := image.NewRGBA(bounds)

	for i := 0; i < bounds.Dx(); i += 1 {
		for j := 0; j < bounds.Dy(); j += 1 {
			var r, g, b, a int = 0, 0, 0, 0

			for k := 0; k < 3; k += 1 {
				for l := 0; l < 3; l += 1 {
					color := img.At(i+(k-1), j+(l-1))
					R, G, B, A := color.RGBA()

					r += int(R / 257)
					g += int(G / 257)
					b += int(B / 257)
					a += int(A / 257)
				}
			}

			color := color.RGBA{uint8(r / 9), uint8(g / 9), uint8(b / 9), uint8(a / 9)}
			m.Set(i, j, color)
		}
	}

	return m
}

func mode(arr []int) (res int) {
	m := make(map[int]int)

	for _, value := range arr {
		if _, ok := m[value]; ok {
			m[value] += 1
		} else {
			m[value] = 0
		}
	}

	for k, v := range m {
		if v > res {
			res = k
		}
	}

	return
}

func LowPassFilterMode(img image.Image) image.Image {
	bounds := img.Bounds()
	m := image.NewRGBA(bounds)

	for i := 0; i < bounds.Dx(); i += 1 {
		for j := 0; j < bounds.Dy(); j += 1 {
			var r []int = []int{}

			for k := 0; k < 3; k += 1 {
				for l := 0; l < 3; l += 1 {
					color := img.At(i+(k-1), j+(l-1))
					R, G, B, _ := color.RGBA()

					r = append(r, int((R+G+B)/(3*257)))
				}
			}

			color := color.RGBA{uint8(mode(r)), uint8(mode(r)), uint8(mode(r)), 255}
			m.Set(i, j, color)
		}
	}

	return m
}

func median(arr []int) int {
	sort.Ints(arr)

	if len(arr) == 0 {
		return 0
	}

	return arr[len(arr)/2]
}

func LowPassFilterMedian(img image.Image) image.Image {
	bounds := img.Bounds()
	m := image.NewRGBA(bounds)

	for i := 0; i < bounds.Dx(); i += 1 {
		for j := 0; j < bounds.Dy(); j += 1 {
			var r []int = []int{}

			for k := 0; k < 3; k += 1 {
				for l := 0; l < 3; l += 1 {
					color := img.At(i+(k-1), j+(l-1))
					R, G, B, _ := color.RGBA()

					r = append(r, int((R+G+B)/(3*257)))
				}
			}

			color := color.RGBA{uint8(median(r)), uint8(median(r)), uint8(median(r)), 255}
			m.Set(i, j, color)
		}
	}

	return m
}
