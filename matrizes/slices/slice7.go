package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var img = make([][]uint8, dx)
	for x := range img {
		img[x] = make([]uint8, dy)
		for y := range img[x] {
			img[x][y] = uint8(y - int(16*math.Sin(float64(x*2))) - 130)
		}
	}
	return img
}
func main() {
	pic.Show(Pic)
}
