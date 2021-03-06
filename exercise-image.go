package main

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/tour/pic"
)

type Point struct {
	X, Y int
}

type Image struct {
	Min, Max Point
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(i.Min.X, i.Min.Y, i.Max.X, i.Max.Y)
}

func (i Image) At(x, y int) color.Color {
	v := uint8(math.Sin(float64((x*x+y*y)/(i.Max.X/4))) * float64(i.Max.X))
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{Point{0, 0}, Point{400, 400}}
	pic.ShowImage(m)
}

