package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), uint8(x - y), 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
