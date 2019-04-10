package main

import "golang.org/x/tour/pic"

type myFunc func(x, y int) int

func (f myFunc) Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8(f(x, y))
		}
	}
	return pic
}

func main() {
	f := myFunc(func(x, y int) int { return x + y })
	pic.Show(f.Pic)
}
