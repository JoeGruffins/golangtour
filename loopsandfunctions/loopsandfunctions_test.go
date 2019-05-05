package main

import (
	"math"
	"testing"
)

type sqrtTestTrio struct {
	input     float64
	outputNum float64
	outputErr error
}

var sqrtTests = []sqrtTestTrio{
	{4, 2, nil},
	{35, math.Sqrt(35), nil},
	{0.1, math.Sqrt(0.1), nil},
	{0, 0, nil},
	{-1, 0, SqrtNegativeError(-1)},
}

func TestSqrt(t *testing.T) {
	for _, trio := range sqrtTests {
		v, err := Sqrt(trio.input)
		if v != trio.outputNum || err != trio.outputErr {
			t.Error("for", trio.input, "expected", trio.outputNum, trio.outputErr, "got", v, err)
		}
	}
}
