package actions

import (
	"fmt"
	"image"
)

type MirrorDirection string

const (
	HORIZONTAL MirrorDirection = "horizontal"
	VERTICAL   MirrorDirection = "vertical"
)

func horizontalFlip(img image.Image) image.Image {
	size := img.Bounds().Size()
	m := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))

	for i := 0; i < size.X; i += 1 {
		for j := 0; j < size.Y; j += 1 {
			m.Set(size.X-i, j, img.At(i, j))
		}
	}

	return m
}

func verticalFlip(img image.Image) image.Image {
	size := img.Bounds().Size()
	m := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))

	for i := 0; i < size.X; i += 1 {
		for j := 0; j < size.Y; j += 1 {
			m.Set(i, size.Y-j, img.At(i, j))
		}
	}

	return m
}

// 0 > factor <= 2
func Mirroring(img image.Image, direction MirrorDirection) image.Image {
	switch direction {
	case HORIZONTAL:
		{
			return horizontalFlip(img)
		}
	case VERTICAL:
		{
			return verticalFlip(img)
		}

	default:
		fmt.Println("got some wrong direction")
	}

	return img
}
