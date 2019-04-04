package main

import (
	"fmt"
)

const DELTA = 0.00000001

var (
	enNegNumString string = `Can't square negative number.`
)

type SqrtNegativeError float64

func (f SqrtNegativeError) Error() string {
	return fmt.Sprintf(enNegNumString+" %v", float64(f))
}

func nextZandDiff(x, z float64) (float64, float64) {
	newZ := z - (z*z-x)/(2*z)
	diff := newZ - z
	if diff < 0 {
		diff *=-1
	}
	return newZ, diff
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtNegativeError(x)
	}
	z := 1.0
	for newZ, diff := nextZandDiff(x, z); diff > DELTA; z = newZ {
		newZ, diff = nextZandDiff(x, z)
	}
	return z, nil
}
