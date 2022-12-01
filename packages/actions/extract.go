package actions

import (
	"fmt"
	"image"
	"image/color"
)

func copyGray(img image.Image) *image.Gray {
	bounds := img.Bounds()

	m := image.NewGray(bounds)

	for x := 0; x < bounds.Dx(); x += 1 {
		for y := 0; y < bounds.Dy(); y += 1 {
			m.Set(x, y, img.At(x, y))

		}
	}

	return m
}

// an object is only valid if we fill a region and it doesn't flood away
type Object struct {
	x, y, x1, y1 int
}

func (o Object) String() string {
	return fmt.Sprint(o.x, o.y, o.x1, o.y1)
}

func floodImage(x, y int, target uint8, newColor color.Gray, m *image.Gray) {
	if m.GrayAt(x, y).Y == target {
		m.SetGray(x, y, newColor)

		// four-way
		floodImage(x+1, y, target, newColor, m)
		floodImage(x-1, y, target, newColor, m)
		floodImage(x, y+1, target, newColor, m)
		floodImage(x, y-1, target, newColor, m)
	}
}

func fillFromCenter(img image.Image) bool {
	bounds := img.Bounds()

	if bounds.Dx() == 0 || bounds.Dy() == 0 {
		return false
	}

	m := copyGray(img)
	var target uint8 = 255 // white
	var newColor = color.Gray{50}

	x, y := bounds.Dx()/2, bounds.Dy()/2 // kinda on center

	// walk over until no pixel is left or have walk around all pixels
	floodImage(x, y, target, newColor, m)

	// walk around borders; if a border has newColor, return false
	for i := 0; i < bounds.Dx(); i += 1 {
		if m.GrayAt(0, i) == newColor {
			return false
		}
	}

	for i := 0; i < bounds.Dx(); i += 1 {
		if m.GrayAt(bounds.Dy(), i) == newColor {
			return false
		}
	}

	for i := 0; i < bounds.Dy(); i += 1 {
		if m.GrayAt(i, 0) == newColor {
			return false
		}
	}
	for i := 0; i < bounds.Dy(); i += 1 {
		if m.GrayAt(i, bounds.Dy()) == newColor {
			return false
		}
	}

	return true
}

func findObjects(img image.Image) (objects []Object) {
	m := copyGray(img)

	if fillFromCenter(m) {
		objects = append(objects, Object{0, 0, m.Bounds().Dx(), m.Bounds().Dy()})
	}

	return
}

func ExtractCharacteristics(img image.Image) (text []string) {
	objects := findObjects(img)

	if len(objects) == 0 {
		return []string{"no object found"}
	}

	for _, obj := range objects {
		text = append(text, obj.String())
	}

	return
}
