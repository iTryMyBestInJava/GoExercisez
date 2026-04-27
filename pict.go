package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		xslice := make([]uint8, dx)

		for x := 0; x < dx; x++ {
			xslice[x] = uint8(x ^ y) //uint8((x+y)/2)
		}

		picture[y] = xslice
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
