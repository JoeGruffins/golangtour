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

type inputTestPair struct {
	input  string
	output bool
}

var sqrtTests = []sqrtTestTrio{
	{4, 2, nil},
	{35, math.Sqrt(35), nil},
	{0.1, math.Sqrt(0.1), nil},
	{0, 0, nil},
	{-1, 0, SqrtNegativeError(-1)},
}

var inputTests = []inputTestPair{
	{"exit", true},
	{"eight", false},
	{"7", false},
	{"0", false},
	{"-3", false},
}

func TestSqrt(t *testing.T) {
	for _, trio := range sqrtTests {
		v, err := Sqrt(trio.input)
		if v != trio.outputNum || err != trio.outputErr {
			t.Error("for", trio.input, "expected", trio.outputNum, trio.outputErr, "got", v, err)
		}
	}
}

func TestInput(t *testing.T) {
	for _, pair := range inputTests {
		v := processInput(pair.input)
		if v != pair.output {
			t.Error("for", pair.input, "expected", pair.output, "got", v)
		}
	}
}
