package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	pixels [][]color.RGBA
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(img.pixels), len(img.pixels[0]))
}

func (img *Image) At(x, y int) color.Color {
	return img.pixels[x][y]
}

const WIDTH = 255
const HEIGHT = 255

func main() {
	img := make([][]color.RGBA, WIDTH)
	for i := 0; i < WIDTH; i++ {
		img[i] = make([]color.RGBA, HEIGHT)
		for j := 0; j < HEIGHT; j++ {
			img[i][j] = color.RGBA{uint8(i * j), uint8(j * i), 255, 255}
		}
	}
	
	m := Image{pixels: img}
	pic.ShowImage(&m)
}
