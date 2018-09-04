package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := 0; i < len(pic); i++ {
		pic[i] = make([]uint8, dx)
		for j := 0; j < len(pic[i]); j++ {
			pic[i][j] = (uint8) (j * i)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
