package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2/container"
)

func About() fyne.CanvasObject {
	c := container.NewVBox()

	c.Add(widget.NewCard("Author", "David Alen", nil))

	return c
}
