package main

import (
	"fmt"
)

const DELTA = 0.00000001

func nextZ(x,z float64) (float64, float64) {
	newZ := z - (z*z - x) / (2*z)
	change := newZ - z
	if change < 0 { change *= -1 }
	return newZ, change
}

func Sqrt(x float64) float64 {
	z := 1.0
	newZ, change := nextZ(x,z)
	for ;change > DELTA; z = newZ{
		newZ, change = nextZ(x,z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(43))
}

