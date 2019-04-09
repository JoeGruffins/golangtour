package main

import "golang.org/x/tour/pic"

type myFunc struct {
	F func(x, y int) int
}

func (g myFunc) Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8(g.F(x, y))
		}
	}
	return pic
}

func main() {
	p := new(myFunc)
	p.F = func(x, y int) int { return x + y }
	pic.Show(p.Pic)
}
