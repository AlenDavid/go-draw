package main

import "github.com/alendavid/go-draw/packages/ui"

func main() {
	if err := ui.Run(); err != nil {
		panic(err)
	}
}
